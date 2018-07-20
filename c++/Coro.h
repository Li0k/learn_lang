//
// Created by liiiyu on 2018/7/17.
//

#ifndef CORO_CORO_H
#define CORO_CORO_H

#include <functional>
//#define _XOPEN_SOURCE 600

#define _XOPEN_SOURCE

#include <ucontext.h>

#undef _XOPEN_SOURCE

#include <unordered_map>
#include <memory>

enum Coro_Satus {
    CORO_DEATH,
    CORO_READY,
    CORO_RUNNING,
    CORO_SUSPEND
};

typedef std::function<void()> coroutine_func;

//#define MAX_STACK_SIZE 1024 * 128
#define DEAFULT_NUM 16


class coroutine {
public:
    int id_;
    ucontext_t ctx;//上下文
    coroutine_func func;//运行函数
    Coro_Satus status;//运行状态
//    char stack[MAX_STACK_SIZE];//运行栈
    char stack_[1024*1024];
//    char* stack_;


    coroutine(const coroutine_func &f,int id);

    ~coroutine();

    Coro_Satus &coroutine_status();
};

typedef std::shared_ptr<coroutine> Coroutine_ptr;

class Coro {
public:
    ucontext_t mctx;//main上下文
//    coroutine *running_co;//正在运行的co
//    std::set<coroutine *> coroutines;//存放co
    int curr;//正在运行的co的id
    int nco;//正在运行的co
    int cap = DEAFULT_NUM;//容量
//    std::vector<coroutine *> coroutines;
    std::unordered_map<int, Coroutine_ptr> coroutines;
    int stack_size;

//    Coro(int ss = MAX_STACK_SIZE);
    Coro();

    ~Coro();

    static void excute_fun(uint32_t low32, uint32_t hi32);

    void resume(int id);

    void yield();

    int create(const coroutine_func &f);

    Coro_Satus status(int id);

    size_t co_num();
};

#endif //CORO_CORO_H
