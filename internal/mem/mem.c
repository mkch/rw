#include <stdlib.h>
#include "mem.h"

unsigned char smallBuffer[SMALL_BUFFER_SIZE] = {};
unsigned char* smallBufferNext = smallBuffer;
