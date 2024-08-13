#import <Foundation/Foundation.h>
#import <CoreServices/CoreServices.h>

func applicationDidFinishLaunching(_ aNotification : Notification)
{
    NSAppleEventManager.shared().setEventHandler(self, andSelector : #selector(self.handleAppleEvent(event:replyEvent:)), forEventClass : AEEventClass(kInternetEventClass), andEventID : AEEventID(kAEGetURL))
}
@objc func handleAppleEvent(event : NSAppleEventDescriptor, replyEvent : NSAppleEventDescriptor)
{
    guard let appleEventDescription = event.paramDescriptor(forKeyword : AEKeyword(keyDirectObject)) else
    {
        return
    }

    if let
        appleEventURLString = appleEventDescription.stringValue
        {
            guard let appleEventURL = URL(string : appleEventURL) else
            {
                return
            }

            if service
                .service_started
                {
                    print(appleEventURL)
                }
        }
}