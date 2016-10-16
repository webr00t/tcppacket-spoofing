package interfaces

//#include <sys/ioctl.h>
//#include <sys/types.h>
//#include <sys/socket.h>
//#ifdef HAVE_LINUX_WIRELESS_H
//# include <linux/wireless.h>
//#else
//# include <net/if.h>
//# include <net/if_media.h>
//#endif
//static int s_wireless_nic (const char* name) {
//    int sock = 0;
//    int result = 0;
//    if ((sock = socket(AF_INET, SOCK_DGRAM, 0)) < 0) {
//        return 0;
//    }
//#   ifdef SIOCGIFMEDIA
//    struct ifmediareq ifmr;
//
//    memset (&ifmr, 0, sizeof (struct ifmediareq));
//    strlcpy(ifmr.ifm_name, name, sizeof(ifmr.ifm_name));
//    if (ioctl(sock, SIOCGIFMEDIA, (caddr_t) &ifmr) != -1) {
//        result = IFM_TYPE (ifmr.ifm_current) == IFM_IEEE80211;
//    }
//#   elif defined(SIOCGIWNAME)
//    struct iwreq wrq;
//
//    strncpy (wrq.ifr_name, name, IFNAMSIZ);
//    if (ioctl(sock, SIOCGIWNAME, (caddr_t) &wrq) != -1) {
//        result = 1;
//    }
//#   endif
//    close(sock);
//    return result;
//}
import "C"

type InterfaceTypeDetector interface {
	IsType(deviceName string) bool
}

type WirelessTypeDetector struct {
}

func (detector WirelessTypeDetector) IsType(deviceName string) bool {
	if C.s_wireless_nic(C.CString(deviceName)) == 1 {
		return true
	}

	return false
}