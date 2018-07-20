#include <iostream>
#include "Coro.h"


coroutine::coroutine(const coroutine_func &f, int id)
        : func(f),
          status(CORO_READY),
          id_(id) {
//    stack_ = new char[MAX_STACK_SIZE];
};
//coroutine::coroutine(const coroutine_func &f)
//        : func(f), status(CORO_READY) {}

Coro_Satus &coroutine::coroutine_status() {
    return status;
}

coroutine::~coroutine() {

}

Coro_Satus Coro::status(int id) {
    assert(id >= 0 && id < cap);
    auto co = coroutines[id];
    if (co == nullptr) {
        return CORO_DEATH;
    }
    return co->status;
}

Coro::Coro()
        : curr(-1) {
}

Coro::~Coro() {

}

void Coro::excute_fun(uint32_t low32, uint32_t hi32) {
    uintptr_t ptr = (uintptr_t) low32 | (uintptr_t(hi32) << 32);
    Coro *schedule = (Coro *) ptr;
    auto id = schedule->curr;
    std::cout << "excute_fun1: " << id << std::endl;
//    Coroutine_ptr co = schedule->coroutines[id];
    auto it = schedule->coroutines.find(id);
    assert(it != schedule->coroutines.end());
    auto co = it->second;
//    std::cout << "excute cap: " << schedule-> << id << std::endl;

    co->func();
    std::cout << "excute_fun 2" << std::endl;
    schedule->coroutines[schedule->curr] = nullptr;
    std::cout << "excute_fun 3" << std::endl;
    co->status = CORO_DEATH;
    schedule->curr = -1;
    std::cout << "excute_fun end" << std::endl;
}

void Coro::resume(int id) {
    std::cout << "resume" << std::endl;
    assert(id >= 0 && id < cap);
    auto it = coroutines.find(id);
    assert(it != coroutines.end());
    assert(curr == -1);

    auto new_co = it->second;
//    assert(new_co->stack_ != nullptr);

    auto status = new_co->status;

    switch (status) {
        case CORO_READY: {
            getcontext(&new_co->ctx);
//            new_co->stack_ = new char[1024 * 1024];
//            memset(new_co->stack_,0,1024*1024);
            new_co->ctx.uc_stack.ss_sp = new_co->stack_;
            new_co->ctx.uc_stack.ss_size = 1024 * 1024;
            new_co->ctx.uc_link = &mctx;
            curr = id;
            new_co->status = CORO_RUNNING;

            uintptr_t ptr = (uintptr_t) this;
            makecontext(&new_co->ctx, (void (*)(void)) excute_fun, 2, (uint32_t) ptr, (uint32_t) (ptr >> 32));
            swapcontext(&mctx, &new_co->ctx);
            break;
        }
        case CORO_SUSPEND: {
            curr = id;
            new_co->status = CORO_RUNNING;
            swapcontext(&mctx, &new_co->ctx);
            break;
        }
        default:
            assert(0);
    }

}

void Coro::yield() {
    assert(curr != -1);
    auto co = coroutines[curr];
    assert(co->status != CORO_DEATH);

    co->status = CORO_SUSPEND;
    curr = -1;
    swapcontext(&co->ctx, &mctx);
}

int Coro::create(const coroutine_func &f) {
    assert(coroutines.size() <= cap);
    int id = -1;

    for (int i = 0; i < cap; ++i) {
        id = (i + nco) % cap;
        auto it = coroutines.find(id);
        if (it == coroutines.end()) {
//            Coroutine_ptr new_co = std::make_shared<coroutine>(new coroutine(f,id));
            Coroutine_ptr new_co(new coroutine(f, id));
//            coroutines[id] = new_co;
            coroutines.insert({id,new_co});
            return id;
        }
    }
    return id;
}

size_t Coro::co_num() {
    return coroutines.size();
}

//
//Coro::Coro(int ss)
//        : stack_size(ss), running_co(nullptr) {
//
//}
