/*
favirecon - Use favicon.ico to improve your target recon phase. Quickly detect technologies, WAF, exposed panels, known services.

This repository is under MIT License https://github.com/edoardottt/favirecon/blob/main/LICENSE
*/

package favirecon

import (
	"errors"
	"fmt"

	"github.com/projectdiscovery/goflags"
)

//nolint: gochecknoglobals
var (
	db = map[string]string{
		"-1010568750": "phpMyAdmin",
		"1015545776":  "pfSense",
		"-1015932800": "Ghost (CMS)",
		"1020814938":  "Ubiquiti - AirOS",
		"-1022206565": "CrushFTP",
		"-1028703177": "TP Link",
		"1037387972":  "Dlink Router",
		"1038500535":  "D-Link (router/network)",
		"-1038557304": "Webmin",
		"-1041180225": "QNAP NAS Virtualization Station",
		"104189364":   "Vigor Router",
		"1045696447":  "Sophos User Portal/VPN Portal",
		"1047213685":  "Netgear (Network)",
		"-1050786453": "Plesk",
		"105083909":   "FireEye",
		"1051648103":  "Securepoint",
		"1064742722":  "RabbitMQ",
		"-106646451":  "WISPR (Airlan)",
		"-1067420240": "GraphQL",
		"1081719753":  "D-Link (Network)",
		"-1085284672": "Wordpress",
		"-1093172228": "truVision (NVR)",
		"1095915848":  "Airwatch",
		"-10974981":   "Shinobi (CCTV)",
		"1103599349":  "Untangle",
		"1117165781":  "SimpleHelp (Remote Support)",
		"1118684072":  "Baidu",
		"-1119613926": "Bluehost",
		"-1124868062": "Netport Software (DSL)",
		"1126835021":  "InstaCart",
		"1135165421":  "Ricoh",
		"1139788073":  "Metasploit",
		"1142227528":  "Aruba (Virtual Controller)",
		"1143877728":  "Dremio",
		"1144925962":  "Dlink Webcam",
		"1147858285":  "Dell",
		"-1148190371": "OPNsense",
		"-1151675028": "ISP Manager (Web Hosting Panel)",
		"-1153873472": "Airwatch",
		"-1153950306": "Dell",
		"1157789622":  "Ubiquiti UNMS",
		"-1160966609": "Gitea",
		"-1162730477": "Vanderbilt SPC",
		"116323821":   "Spring boot",
		"1169183049":  "BoaServer",
		"-1169314298": "INSTAR IP Cameras",
		"1174841451":  "Drupal",
		"1179099333":  "Youtube",
		"1188645141":  "Huweishen",
		"119741608":   "Teltonika",
		"12003995":    "Walmart",
		"-1200737715": "Kibana",
		"-1203021870": "Kubeflow",
		"-1205024243": "lwIP (A Lightweight TCP/IP stack)",
		"1211608009":  "Openfire Admin Console",
		"1221759509":  "Dlink Webcam",
		"-1225484776": "Endian Firewall",
		"1227052603":  "Alibaba Cloud (Block Page)",
		"-1231681737": "Ghost (CMS)",
		"1232159009":  "Apple",
		"1232596212":  "OpenStack",
		"1234311970":  "Onera",
		"1235070469":  "Synology VPN Plus",
		"-1235192469": "Metasploit",
		"123821839":   "Sangfor",
		"-1238837624": "Liquid Pixels",
		"-1240222446": "Zhejiang Uniview Technologies Co.",
		"1241049726":  "iomega NAS",
		"1244636413":  "cPanel Login",
		"1248917303":  "JupyterHub",
		"1249285083":  "Ubiquiti Aircube",
		"-1249852061": "Microsoft Outlook",
		"1251810433":  "Cafe24 (Korea)",
		"-1252041730": "Vue.js",
		"-1255347784": "AngularJS",
		"-1255992602": "VMware Horizon",
		"1262005940":  "Jamf Pro Login",
		"-1267819858": "KeyHelp (Keyweb AG)",
		"-1268095485": "VZPP Plesk",
		"-12700016":   "Seafile",
		"1273982002":  "Mautic (Open Source Marketing Automation)",
		"1274078387":  "TP-LINK (Network Device)",
		"-1275148624": "Accrisoft",
		"-1275226814": "XAMPP",
		"-1277814690": "LaCie",
		"1278323681":  "Gitlab",
		"-127886975":  "Metasploit",
		"1280907310":  "Webmin",
		"1281253102":  "Dahua Storm (DVR)",
		"129457226":   "Liferay Portal",
		"-1298108480": "Yii PHP Framework (Default Favicon)",
		"1302486561":  "NetData",
		"1307375944":  "Octoprint (3D printer)",
		"1318124267":  "Avigilon",
		"-1319025408": "Netgear",
		"1319699698":  "Form.io",
		"1333537166":  "Alfresco",
		"-1343070146": "Intelbras SA",
		"-134375033":  "Plesk",
		"-1346447358": "TilginAB (HomeGateway)",
		"1347937389":  "SAP Conversational AI",
		"-1351901211": "Luma Surveillance",
		"-1354933624": "Dlink Webcam",
		"1356662359":  "Outlook Web Application",
		"1370528867":  "Yahoo",
		"-137295400":  "NETGEAR ReadyNAS",
		"-1378182799": "Archivematica",
		"-1379982221": "Atlassian - Bamboo",
		"1382324298":  "Apple",
		"-1395400951": "Huawei - ADSL/Router",
		"-1399433489": "Prometheus Time Series Collection and Processing Server",
		"1405460984":  "pfSense",
		"1410610129":  "Supermicro Intelligent Management (IPMI)",
		"-1414475558": "Microsoft IIS",
		"142313513":   "Facebook",
		"-1424036600": "Portainer",
		"1424295654":  "Icecast Streaming Media Server",
		"1427976651":  "ZTE (Network)",
		"1433417005":  "Salesforce",
		"1436966696":  "Barracuda",
		"-1437701105": "XAMPP",
		"-1441956789": "Tableau",
		"-1442789563": "Nuxt JS",
		"-1446794564": "Ubiquiti Login Portals",
		"-1452159623": "Tecvoz",
		"1453890729":  "Webmin",
		"-1457536113": "CradlePoint",
		"-1457628171": "Postmark",
		"1462981117":  "Cyberoam",
		"-1465479343": "DNN (CMS)",
		"-1466785234": "Dahua",
		"1466912879":  "CradlePoint Technology (Router)",
		"1467395679":  "Ligowave (network)",
		"-1474875778": "GLPI",
		"1476335317":  "FireEye",
		"-1477563858": "Arris",
		"1479202414":  "Arcadyan o2 box (Network)",
		"1483097076":  "SyncThru Web Service (Printers)",
		"1485257654":  "SonarQube",
		"1490343308":  "MK-AUTH",
		"-1492966240": "RADIX",
		"149371702":   "Synology DiskStation",
		"-1498185948": "Apple",
		"1502482995":  "Indeed",
		"-1507567067": "Baidu (IP error page)",
		"1528355650":  "Pinterest",
		"-1528414776": "Rumpus",
		"1537743656":  "DropBox",
		"1540037626":  "VKontakte",
		"1544230796":  "cPanel Login",
		"-1544605732": "Amazon",
		"-1546574541": "Sonatype Nexus Repository Manager",
		"-1547576879": "Saia Burgess Controls - PCD",
		"1552860581":  "Elastic (Database)",
		"-1561873722": "Nginx",
		"156312019":   "Technicolor / Thomson Speedtouch (Network / ADSL)",
		"-1571472432": "Sierra Wireless Ace Manager (Airlink)",
		"-1581907337": "Atlassian - JIRA",
		"15831193":    "WatchGuard",
		"1585145626":  "netdata dashboard",
		"-1588746893": "CommuniGate",
		"-1589842876": "Deluge Web UI",
		"-1593651747": "Blackboard",
		"1594377337":  "Technicolor",
		"1601194732":  "Sophos Cyberoam (appliance)",
		"1603223646":  "Comcast Business",
		"-160425702":  "Medallia",
		"-1607644090": "Bitnami",
		"1611729805":  "Elastic (Database)",
		"-1612496354": "Teltonika",
		"-1616115760": "ownCloud",
		"-1616143106": "AXIS (network cameras)",
		"16202868":    "Universal Devices (UD)",
		"1627330242":  "Joomla",
		"1629518721":  "macOS Server (Apple)",
		"-1630354993": "Proofpoint",
		"1632780968":  "Université Toulouse 1 Capitole",
		"163842882":   "Cisco Meraki",
		"-1640178715": "Reddit",
		"-1642532491": "Atlassian - Confluence",
		"1642701741":  "Vmware Secure File Transfer",
		"1648531157":  "InfiNet Wireless | WANFleX (Network)",
		"-1654229048": "Vivotek (Camera)",
		"-1656695885": "iomega NAS",
		"165976831":   "Vodafone (Technicolor)",
		"-166151761":  "Abilis (Network/Automation)",
		"-1666561833": "Wildfly",
		"1668183286":  "Kibana",
		"1673203892":  "Oracle",
		"-167656799":  "Drupal",
		"-1677255344": "UBNT Router UI",
		"1678170702":  "Asustor",
		"-1678298769": "Kerio Connect WebMail",
		"-1688698891": "SpamExperts",
		"-1697334194": "Univention Portal",
		"-1702393021": "mofinetwork",
		"-1702769256": "Bosch Security Systems (Camera)",
		"1703788174":  "D-Link (router/network)",
		"-1710631084": "Askey Cable Modem",
		"-1723752240": "Microhard Systems",
		"1726027799":  "IBM Server",
		"1732786188":  "Apache",
		"-1734573358": "TC-Group",
		"1734609466":  "JustHost",
		"1735289686":  "Whatsapp",
		"-1738184811": "cacaoweb",
		"-1738727418": "KeepItSafe Management Console",
		"-1745552996": "Arbor Networks",
		"-1748763891": "INSTAR Full-HD IP-Camera",
		"-175283071":  "Dell",
		"1768726119":  "Outlook Web Application",
		"1770799630":  "bintec elmeg",
		"1772087922":  "ASP.net",
		"-1775553655": "Unified Management Console (Polycom)",
		"-1779611449": "Alienvault",
		"1782271534":  "truVision NVR (interlogix)",
		"1786752597":  "wdCP cloud host management system",
		"-178685903":  "Yasni",
		"-1788112745": "PowerMTA monitoring",
		"1802374283":  "LiquidPixels",
		"-1807411396": "Skype",
		"-1810847295": "Sangfor",
		"-1814887000": "Docker",
		"1821549811":  "(Blank) iSpy",
		"-1822098181": "Checkpoint (Gaia)",
		"-182423204":  "netdata dashboard",
		"-1831547740": "Twitter",
		"-183163807":  "Ace",
		"1835479497":  "Technicolor Gateway",
		"1836828108":  "OpenProject",
		"1838417872":  "Freebox OS",
		"-1840324437": "Microsoft 365",
		"1842351293":  "TP-LINK (Network Device)",
		"1848946384":  "GitHub",
		"1862132268":  "Gargoyle Router Management Utility",
		"-1863663974": "Airwatch",
		"1876585825":  "ALIBI NVR",
		"1877797890":  "Eltex (Router)",
		"1895360511":  "VMware Horizon",
		"-1897829998": "D-Link (camera)",
		"1911253822":  "UPC Ceska Republica (Network)",
		"1913538826":  "Material Dashboard",
		"1914658187":  "CloudFlare",
		"191654058":   "Wordpress Under Construction Icon",
		"1922032523":  "NEC WebPro",
		"-1922044295": "Mitel Networks (MiCollab End User Portal)",
		"-1926484046": "Sangfor",
		"-1929912510": "NETASQ - Secure / Stormshield",
		"-1933493443": "Residential Gateway",
		"-1935525788": "SmarterMail",
		"1937209448":  "Docker",
		"1941381095":  "openWRT Luci",
		"1941681276":  "Amazon",
		"-1944119648": "TeamCity",
		"-194439630":  "Avtech IP Surveillance (Camera)",
		"-1950415971": "Joomla",
		"1953726032":  "Metabase",
		"1954835352":  "Vesta Hosting Control Panel",
		"-195508437":  "iPECS",
		"-1961046099": "Dgraph Ratel",
		"1966198264":  "OpenERP (now known as Odoo)",
		"1969970750":  "Gitea",
		"1973665246":  "Entrolink",
		"1975413433":  "Sunny WebBox",
		"1985721423":  "WorldClient for Mdaemon",
		"1991136554":  "Instagram",
		"1991562061":  "Niagara Web Server / Tridium",
		"1993518473":  "cPanel Login",
		"-2006308185": "OTRS (Open Ticket Request System)",
		"2006716043":  "Intelbras SA",
		"-2009722838": "React",
		"2019488876":  "Dahua Storm (IP Camera)",
		"-2031183903": "D-Link (Network)",
		"-2042679530": "Alibaba",
		"2047156994":  "Linksys",
		"-2054889066": "Sentora",
		"2055322029":  "Realtek",
		"-2056503929": "iomega NAS",
		"2059618623":  "HP iLO",
		"-2063036701": "Linksys Smart Wi-Fi",
		"2063428236":  "Sentry",
		"2068154487":  "Digium (Switchvox)",
		"-2069844696": "Ruckus Wireless",
		"2071993228":  "Nomadix Access Gateway",
		"2072198544":  "Ferozo Panel",
		"2086228042":  "MobileIron",
		"2099342476":  "PKP (OpenJournalSystems) Public Knowledge Project",
		"2109473187":  "Huawei - Claro",
		"2113497004":  "WiJungle",
		"-2116540786": "bet365",
		"-2117390767": "Spiceworks (panel)",
		"2119159060":  "GMail",
		"2121539357":  "FireEye",
		"2124459909":  "HFS (HTTP File Server)",
		"-2125083197": "Windows Azure",
		"2127152956":  "MailWizz",
		"2128230701":  "Chainpoint",
		"-2133341160": "WordPress Org",
		"-2138771289": "Technicolor",
		"-2140379067": "RoundCube Webmail",
		"2141724739":  "Juniper Device Manager",
		"-2144363468": "HP Printer / Server",
		"-2145085239": "Tenda Web Master",
		"2146763496":  "Mailcow",
		"-219752612":  "FRITZ!Box",
		"-222497010":  "JoyRun",
		"224536051":   "Shenzhen coship electronics co.",
		"225632504":   "Rocket Chat",
		"-235701012":  "Cnservers LLC",
		"239966418":   "Microsoft Outlook",
		"240136437":   "Seagate Technology (NAS)",
		"240606739":   "FireEye",
		"246145559":   "Parse",
		"251106693":   "GPON Home Gateway",
		"-254193850":  "React",
		"252728887":   "DD WRT (DD-WRT milli_httpd)",
		"255892555":   "wdCP cloud host management system",
		"-256828986":  "iDirect Canada (Network Management)",
		"-266008933":  "SAP Netweaver",
		"-267431135":  "Kibana",
		"-271448102":  "iKuai Networks",
		"-276759139":  "Chef Automate",
		"-277464596":  "AEM Screens",
		"281559989":   "Huawei",
		"283740897":   "Intelbras SA",
		"29056450":    "Geneko",
		"-291579889":  "WS server test page",
		"-297069493":  "Apache Tomcat",
		"-299287097":  "Cisco Router",
		"-299324825":  "Lupus Electronics XT",
		"-305179312":  "Atlassian - Confluence",
		"309020573":   "PayPal",
		"314969666":   "Amazon AWS",
		"-318947884":  "Palo Alto Networks",
		"-318968846":  "ngX-Rocket",
		"31972968":    "Dlink Webcam",
		"321591353":   "Node-RED",
		"321909464":   "Airwatch",
		"322531336":   "iomega NAS",
		"-325082670":  "LANCOM Systems",
		"-329747115":  "C-Lodop",
		"331870709":   "iomega NAS",
		"-332324409":  "STARFACE VoIP Software",
		"-333791179":  "Adobe Campaign Classic",
		"-335153896":  "Traccar GPS tracking",
		"-335242539":  "f5 Big IP",
		"-336242473":  "Siemens OZW772",
		"-342262483":  "Combivox",
		"-35107086":   "WorldClient for Mdaemon",
		"-355305208":  "D-Link (camera)",
		"-359621743":  "Intelbras Wireless",
		"-360566773":  "ARRIS (Network)",
		"362091310":   "MobileIron",
		"363324987":   "Dell SonicWALL",
		"366524387":   "Joomla",
		"-368490461":  "Entronix Energy Management Platform",
		"-373674173":  "Digital Keystone (DK)",
		"-374235895":  "Ossia (Provision SR) | Webcam/IP Camera",
		"-375623619":  "bintec elmeg",
		"381100274":   "Moxapass ioLogik Remote Ethernet I/O Server",
		"-38580010":   "Magento",
		"-386189083":  "aaPanel",
		"-38705358":   "Reolink",
		"-393788031":  "Flussonic (Video Streaming)",
		"396533629":   "OpenVPN",
		"-398568076":  "Wikipedia",
		"-401934945":  "iomega NAS",
		"420473080":   "Exostar - Managed Access Gateway",
		"-421986013":  "Homegrown Website Hosting",
		"-429287806":  "Ebay",
		"430582574":   "SmartPing",
		"-43161126":   "Slack",
		"99395752":    "Slack,",
		"432733105":   "Pi Star",
		"-435817905":  "Cambium Networks",
		"-438482901":  "Moodle",
		"-440644339":  "Zyxel ZyWALL",
		"442749392":   "Microsoft OWA",
		"443944613":   "WAMPSERVER",
		"-450254253":  "idera",
		"-459291760":  "Workday",
		"459900502":   "ZTE Corporation (Gateway/Appliance)",
		"462223993":   "Jeedom (home automation)",
		"-466504476":  "Kerio Control Firewall",
		"475379699":   "Axcient Replibit Management Server",
		"476213314":   "Exacq",
		"-476231906":  "phpMyAdmin",
		"479413330":   "Webmin",
		"483383992":   "ISPConfig",
		"-484708885":  "Zyxel ZyWALL",
		"494866796":   "Aplikasi",
		"-50306417":   "Kyocera (Printer)",
		"-505448917":  "Discuz!",
		"509789953":   "Farming Simulator Dedicated Server",
		"512590457":   "Trendnet IP camera",
		"516963061":   "Gitlab",
		"517158172":   "D-Link (router/network)",
		"-519765377":  "Parallels Plesk Panel",
		"-520888198":  "Blue Iris (Webcam)",
		"-532394952":  "CX",
		"538585915":   "Lenel",
		"541087742":   "LiquidFiles",
		"545827989":   "MobileIron",
		"547025948":   "Grafana",
		"5471989":     "Netcom Technology",
		"547282364":   "Keenetic",
		"547474373":   "TOTOLINK (network)",
		"552592949":   "ASUS AiCloud",
		"552597979":   "Sails",
		"552727997":   "Atlassian - JIRA",
		"5542029":     "NetComWireless (Network)",
		"-560297467":  "DVR (Korean)",
		"56079838":    "Okta",
		"-569941107":  "Fireware Watchguard",
		"575613323":   "Canvas LMS (Learning Management)",
		"577446824":   "Bluehost",
		"579239725":   "Metasploit",
		"586998417":   "Nginx",
		"-587741716":  "ADB Broadband (Network)",
		"-590892202":  "Surfilter SSL VPN Portal",
		"593396886":   "StackOverflow",
		"-594256627":  "Netis (network devices)",
		"-600508822":  "iomega NAS",
		"602431586":   "Palo Alto Login Portal",
		"603314":      "Redmine",
		"607846949":   "Hitron Technologies",
		"-609520537":  "OpenGeo Suite",
		"-613216179":  "iomega NAS",
		"-617743584":  "Odoo",
		"-624805968":  "Cloudinary",
		"-625364318":  "OkoFEN Pellematic",
		"628535358":   "Atlassian",
		"-629047854":  "Jetty 404",
		"-630493013":  "DokuWiki",
		"-631002664":  "Kerio Control Firewall",
		"631108382":   "SonicWALL",
		"-632070065":  "Apache Haus",
		"-632583950":  "Shoutcast Server",
		"-644617577":  "SmartLAN/G",
		"-649378830":  "WHM",
		"-652508439":  "Parallels Plesk Panel",
		"-655683626":  "PRTG Network Monitor",
		"-656811182":  "Jboss",
		"656868270":   "iomega NAS",
		"661332347":   "MOBOTIX Camera",
		"671221099":   "innovaphone",
		"-675839242":  "openWRT Luci",
		"-676077969":  "Niagara Web Server",
		"-677167908":  "Kerio Connect (Webmail)",
		"-687783882":  "ClaimTime (Ramsell Public Health & Safety)",
		"-689902428":  "iomega NAS",
		"-692947551":  "Ruijie Networks (Login)",
		"-693082538":  "openmediavault (NAS)",
		"-696586294":  "LinkedIn",
		"693122507":   "WordPress",
		"-697231354":  "Ubiquiti - AirOS",
		"-702384832":  "TCN",
		"705143395":   "Atlassian",
		"706602230":   "VisualSVN Server",
		"708578229":   "Google",
		"711742418":   "Hitron Technologies Inc.",
		"716989053":   "Amazon AWS",
		"72005642":    "RemObjects SDK / Remoting SDK for .NET HTTP Server Microsoft",
		"-723685921":  "Oracle Cloud",
		"726817668":   "KeyHelp (Keyweb AG)",
		"727253975":   "Paradox IP Module",
		"728788645":   "IBM Notes",
		"731374291":   "HFS (HTTP File Server)",
		"-736276076":  "MyASP",
		"-740211187":  "Bing",
		"743365239":   "Atlassian",
		"74935566":    "WindRiver-WebServer",
		"75230260":    "Kibana",
		"758890177":   "Tumblr",
		"-759108386":  "Tongda",
		"-759754862":  "Kibana",
		"76658403":    "TheTradeDesk",
		"-766957661":  "MDaemon Webmail",
		"768231242":   "JAWS Web Server (IP Camera)",
		"768816037":   "UniFi Video Controller (airVision)",
		"77044418":    "Polycom",
		"-771764544":  "Parallels Plesk Panel",
		"774252049":   "FastPanel Hosting",
		"784872924":   "Lucee!",
		"786476039":   "AppsFlyer",
		"786533217":   "OpenStack",
		"788771792":   "Airwatch",
		"794809961":   "CheckPoint",
		"804949239":   "Cisco Meraki Dashboard",
		"812385209":   "Solarwinds Serv-U FTP Server",
		"81586312":    "Jenkins",
		"-816821232":  "GitLab",
		"829321644":   "BOMGAR Support Portal",
		"-831826827":  "NOS Router",
		"833190513":   "Dahua Storm (IP Camera)",
		"-842192932":  "FireEye",
		"855273746":   "JIRA",
		"86919334":    "ServiceNow",
		"-873627015":  "HeroSpeed Digital Technology Co. (NVR/IPC/XVR)",
		"878647854":   "BIG-IP",
		"-878891718":  "Twonky Server (Media Streaming)",
		"882208493":   "Lantronix (Spider)",
		"-882760066":  "ZyXEL (Network)",
		"-884776764":  "Huawei (Network)",
		"892542951":   "Zabbix",
		"89321398":    "Askey Cable Modem",
		"-895890586":  "PLEX Server",
		"-895963602":  "Jupyter Notebook",
		"896412703":   "IW",
		"899457975":   "Cisco",
		"90066852":    "JAWS Web Server (IP Camera)",
		"902521196":   "Netflix",
		"903086190":   "Honeywell",
		"904434662":   "Loxone (Automation)",
		"905744673":   "HP Printer / Server",
		"905796143":   "Medallia",
		"90680708":    "Domoticz (Home Automation)",
		"916642917":   "Multilaser",
		"917966895":   "Gogs",
		"920338972":   "Linode",
		"-923088984":  "OpenStack",
		"-923693877":  "motionEye (camera)",
		"926501571":   "Handle Proxy",
		"929825723":   "WAMPSERVER",
		"936297245":   "Twitch",
		"937999361":   "JBoss Application Server 7",
		"938616453":   "Mersive Solstice",
		"943925975":   "ZyXEL",
		"944969688":   "Deluge",
		"945408572":   "Fortinet - Forticlient",
		"95271369":    "FireEye",
		"955369722":   "Sitecore",
		"-956471263":  "Web Client Pro",
		"966563415":   "WordPress Org",
		"967636089":   "MobileIron",
		"970132176":   "3CX Phone System",
		"-972810761":  "HostMonster - Web hosting",
		"97604680":    "Tandberg",
		"-976235259":  "Roundcube Webmail",
		"-978656757":  "NETIASPOT (Network)",
		"979634648":   "StruxureWare (Schneider Electric)",
		"980692677":   "Cake PHP",
		"-981606721":  "Plesk",
		"981867722":   "Atlassian - JIRA",
		"-986678507":  "ISP Manager",
		"-986816620":  "OpenRG",
		"987967490":   "Huawei (Network)",
		"988422585":   "CapRover",
		"-991123252":  "VMware Horizon",
		"99432374":    "MDaemon Remote Administration",
		"998138196":   "iomega NAS",
		"999357577":   "Hikvision camera",
	}

	ErrHashNotFound    = errors.New("hash not found")
	ErrHashNotMatching = errors.New("hash not matching hash provided")
)

func CheckFavicon(faviconHash string, hash goflags.StringSlice, url ...string) (string, error) {
	if k, ok := db[faviconHash]; ok {
		if len(hash) != 0 {
			if contains(hash, faviconHash) {
				return k, nil
			}

			return "", fmt.Errorf("[%s] %s %w", faviconHash, url, ErrHashNotMatching)
		}

		return k, nil
	}

	if len(url) == 0 {
		return "", fmt.Errorf("%w", ErrHashNotFound)
	}

	return "", fmt.Errorf("[%s] %s %w", faviconHash, url, ErrHashNotFound)
}
