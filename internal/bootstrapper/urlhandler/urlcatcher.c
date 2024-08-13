#include <stdio.h>
#include <CoreServices/CoreServices.h>

int set(int argc, char *argv[])
{
	if (argc != 4)
	{
		fprintf(stderr, "Usage: %s <set/check> <URL_SCHEME> <BUNDLE_IDENTIFIER>\n", argv[0]);
		return 1;
	}

	const char *command = argv[1];
	const char *urlScheme = argv[2];
	const char *bundleIdentifier = argv[3];

	if (strcmp(command, "set") == 0)
	{
		OSStatus status = LSSetDefaultHandlerForURLScheme(CFStringCreateWithCString(NULL, urlScheme, kCFStringEncodingUTF8),
														  CFStringCreateWithCString(NULL, bundleIdentifier, kCFStringEncodingUTF8));

		if (status == noErr)
		{
			printf("Successfully set %s as the default handler for %s://\n", bundleIdentifier, urlScheme);
			return 0;
		}
		else
		{
			fprintf(stderr, "Failed to set default handler. Error code: %d\n", (int)status);
			return 1;
		}
	}
	else if (strcmp(command, "check") == 0)
	{
		CFStringRef currentHandler = LSCopyDefaultHandlerForURLScheme(CFStringCreateWithCString(NULL, urlScheme, kCFStringEncodingUTF8));

		if (currentHandler)
		{
			Boolean isDefault = CFStringCompare(currentHandler, CFStringCreateWithCString(NULL, bundleIdentifier, kCFStringEncodingUTF8), 0) == kCFCompareEqualTo;
			CFRelease(currentHandler);

			if (isDefault)
			{
				printf("true\n");
				return 0;
			}
			else
			{
				printf("false\n");
				return 1;
			}
		}
		else
		{
			fprintf(stderr, "No default handler found for %s://\n", urlScheme);
			return 1;
		}
	}
	else
	{
		fprintf(stderr, "Invalid command. Use 'set' to set the handler or 'check' to check the handler.\n");
		return 1;
	}
}

int main()
{
	// set roblox scheme too later
	char *argv[] = {"main", "check", "roblox-player", "com.roblox.RobloxPlayer"};
	set(4, argv);

	return 0;
}
