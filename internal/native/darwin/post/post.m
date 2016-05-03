#import <Foundation/NSThread.h>
#import <Foundation/NSValue.h>
#import "post.h"

static id g_safePoster = nil;
static id g_unsafePoster = nil;


@interface Poster : NSObject {
	FN_POST_CALLBACK callback;
}
@end

@implementation Poster

- (void)run:(NSNumber*)userData {
	callback(userData.unsignedLongValue);
}

- (instancetype)initWithCallback:(FN_POST_CALLBACK)aCallback {
	self = [super init];
	if(self) {
		callback = aCallback;
	}
	return self;
}

@end


void initPostOnMainThread(FN_POST_CALLBACK safeCallback, FN_POST_CALLBACK unsafeCallback) {
	g_safePoster = [[Poster alloc] initWithCallback:safeCallback];
	g_unsafePoster = [[Poster alloc] initWithCallback:unsafeCallback];
}

void postOnMainThread(UINTPTR userData, bool safe) {
	[safe?g_safePoster:g_unsafePoster performSelectorOnMainThread:@selector(run:)
							   withObject:[NSNumber numberWithUnsignedLong:userData]
							waitUntilDone:NO];
}