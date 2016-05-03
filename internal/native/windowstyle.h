#ifndef _RW_WINDOWSTYLE_COMMON_H
#define _RW_WINDOWSTYLE_COMMON_H

struct WindowStyleFeatures {
	bool hasBorder;
	bool hasTitle;
	bool hasCloseButton;
	bool hasMinimizeButton;
	bool hasMaximizeButton;
	bool resizable;
};

typedef struct WindowStyleFeatures WindowStyleFeatures;

#endif