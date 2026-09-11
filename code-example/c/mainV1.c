#include <stdio.h>
#include <threads.h>

int say_hello(void *args) {
    const char* name = (const char*) args;
    printf("Hello %s!\n", name);
    return 0;
}

int main() {
    thrd_t thread;

    const char* name = "Brian";

    int creationReturnCode = thrd_create(&thread, say_hello, (void*)name);
    if (creationReturnCode != thrd_success) {
        fprintf(stderr, "An error happened when creating the thread.");
        return 1;
    }

    int returnByTheFunction;
    // Blocking
    thrd_join(thread, &returnByTheFunction);

    printf("The function called by the thread returned with a code of %d\n", returnByTheFunction);

    return 0;
}
