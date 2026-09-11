#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <threads.h>

typedef struct {
    thrd_t* thread;
    thrd_t* thread_result;
    char* name;
    int result;
} ThreadInfo;

int functionToExecute(void *args) {
    const ThreadInfo* thread = (const ThreadInfo*) args;
    printf("Hello from the thread: %s!\n", thread->name);
    return thread->result;
}

int thread_result(void *args){
    const ThreadInfo* thread = (const ThreadInfo*) args;

    int result;
    // Blocking
    thrd_join(*thread->thread, &result);
    printf("The thread %s, has finished with result %d\n", thread->name, result);
    return 0;
}

void launch_thread(ThreadInfo *threadInfo) {
    if (thrd_create(threadInfo->thread, functionToExecute, threadInfo) != thrd_success) {
        fprintf(stderr, "Error creating the thread %s\n", threadInfo->name);
    }

    if (thrd_create(threadInfo->thread_result , thread_result, threadInfo) != thrd_success) {
        fprintf(stderr, "Error creating the result thread for %s\n", threadInfo->name);
    }
}

int main(){
    thrd_t thread1;
    thrd_t thread2;

    thrd_t thread1Result;
    thrd_t thread2Result;


    ThreadInfo* threadInfo1 = malloc(sizeof(ThreadInfo));
    threadInfo1->thread = &thread1;
    threadInfo1->thread_result = &thread1Result;
    threadInfo1->name = "thread number 1";
    threadInfo1->result = 0;

    ThreadInfo* threadInfo2 = malloc(sizeof(ThreadInfo));
    threadInfo2->thread = &thread2;
    threadInfo2->thread_result = &thread2Result;
    threadInfo2->name = "thread number 2";
    threadInfo2->result = 1;

    printf("Starting threads\n");

    launch_thread(threadInfo1);
    launch_thread(threadInfo2);

    printf("Ending threads\n");

    sleep(1);

    return 0;
}
