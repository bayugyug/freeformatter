package main

import ()

var jsonStr = []byte(`[
{
"Name": "Andrew Toolkit",
"Mime": "application\/andrew-inset",
"Ext": "N\/A"
},
{
"Name": "Applixware",
"Mime": "application\/applixware",
"Ext": ".aw"
},
{
"Name": "Atom Syndication Format",
"Mime": "application\/atom+xml",
"Ext": ".atom, .xml"
},
{
"Name": "Atom Publishing Protocol",
"Mime": "application\/atomcat+xml",
"Ext": ".atomcat"
},
{
"Name": "Atom Publishing Protocol Service Document",
"Mime": "application\/atomsvc+xml",
"Ext": ".atomsvc"
},
{
"Name": "Voice Browser Call Control",
"Mime": "application\/ccxml+xml,",
"Ext": ".ccxml"
},
{
"Name": "Cloud Data Management Interface (CDMI) - Capability",
"Mime": "application\/cdmi-capability",
"Ext": ".cdmia"
},
{
"Name": "Cloud Data Management Interface (CDMI) - Contaimer",
"Mime": "application\/cdmi-container",
"Ext": ".cdmic"
},
{
"Name": "Cloud Data Management Interface (CDMI) - Domain",
"Mime": "application\/cdmi-domain",
"Ext": ".cdmid"
},
{
"Name": "Cloud Data Management Interface (CDMI) - Object",
"Mime": "application\/cdmi-object",
"Ext": ".cdmio"
},
{
"Name": "Cloud Data Management Interface (CDMI) - Queue",
"Mime": "application\/cdmi-queue",
"Ext": ".cdmiq"
},
{
"Name": "CU-SeeMe",
"Mime": "application\/cu-seeme",
"Ext": ".cu"
},
{
"Name": "Web Distributed Authoring and Versioning",
"Mime": "application\/davmount+xml",
"Ext": ".davmount"
},
{
"Name": "Data Structure for the Security Suitability of Cryptographic Algorithms",
"Mime": "application\/dssc+der",
"Ext": ".dssc"
},
{
"Name": "Data Structure for the Security Suitability of Cryptographic Algorithms",
"Mime": "application\/dssc+xml",
"Ext": ".xdssc"
},
{
"Name": "ECMAScript",
"Mime": "application\/ecmascript",
"Ext": ".es"
},
{
"Name": "Extensible MultiModal Annotation",
"Mime": "application\/emma+xml",
"Ext": ".emma"
},
{
"Name": "Electronic Publication",
"Mime": "application\/epub+zip",
"Ext": ".epub"
},
{
"Name": "Efficient XML Interchange",
"Mime": "application\/exi",
"Ext": ".exi"
},
{
"Name": "Portable Font Resource",
"Mime": "application\/font-tdpfr",
"Ext": ".pfr"
},
{
"Name": "Hyperstudio",
"Mime": "application\/hyperstudio",
"Ext": ".stk"
},
{
"Name": "Internet Protocol Flow Information Export",
"Mime": "application\/ipfix",
"Ext": ".ipfix"
},
{
"Name": "Java Archive",
"Mime": "application\/java-archive",
"Ext": ".jar"
},
{
"Name": "Java Serialized Object",
"Mime": "application\/java-serialized-object",
"Ext": ".ser"
},
{
"Name": "Java Bytecode File",
"Mime": "application\/java-vm",
"Ext": ".class"
},
{
"Name": "JavaScript",
"Mime": "application\/javascript",
"Ext": ".js"
},
{
"Name": "JavaScript Object Notation (JSON)",
"Mime": "application\/json",
"Ext": ".json"
},
{
"Name": "Macintosh BinHex 4.0",
"Mime": "application\/mac-binhex40",
"Ext": ".hqx"
},
{
"Name": "Compact Pro",
"Mime": "application\/mac-compactpro",
"Ext": ".cpt"
},
{
"Name": "Metadata Authority  Description Schema",
"Mime": "application\/mads+xml",
"Ext": ".mads"
},
{
"Name": "MARC Formats",
"Mime": "application\/marc",
"Ext": ".mrc"
},
{
"Name": "MARC21 XML Schema",
"Mime": "application\/marcxml+xml",
"Ext": ".mrcx"
},
{
"Name": "Mathematica Notebooks",
"Mime": "application\/mathematica",
"Ext": ".ma"
},
{
"Name": "Mathematical Markup Language",
"Mime": "application\/mathml+xml",
"Ext": ".mathml"
},
{
"Name": "Mbox database files",
"Mime": "application\/mbox",
"Ext": ".mbox"
},
{
"Name": "Media Server Control Markup Language",
"Mime": "application\/mediaservercontrol+xml",
"Ext": ".mscml"
},
{
"Name": "Metalink",
"Mime": "application\/metalink4+xml",
"Ext": ".meta4"
},
{
"Name": "Metadata Encoding and Transmission Standard",
"Mime": "application\/mets+xml",
"Ext": ".mets"
},
{
"Name": "Metadata Object Description Schema",
"Mime": "application\/mods+xml",
"Ext": ".mods"
},
{
"Name": "MPEG-21",
"Mime": "application\/mp21",
"Ext": ".m21"
},
{
"Name": "MPEG4",
"Mime": "application\/mp4",
"Ext": ".mp4"
},
{
"Name": "Microsoft Word",
"Mime": "application\/msword",
"Ext": ".doc"
},
{
"Name": "Material Exchange Format",
"Mime": "application\/mxf",
"Ext": ".mxf"
},
{
"Name": "Binary Data",
"Mime": "application\/octet-stream",
"Ext": ".bin"
},
{
"Name": "Office Document Architecture",
"Mime": "application\/oda",
"Ext": ".oda"
},
{
"Name": "Open eBook Publication Structure",
"Mime": "application\/oebps-package+xml",
"Ext": ".opf"
},
{
"Name": "Ogg",
"Mime": "application\/ogg",
"Ext": ".ogx"
},
{
"Name": "Microsoft OneNote",
"Mime": "application\/onenote",
"Ext": ".onetoc"
},
{
"Name": "XML Patch Framework",
"Mime": "application\/patch-ops-error+xml",
"Ext": ".xer"
},
{
"Name": "Adobe Portable Document Format",
"Mime": "application\/pdf",
"Ext": ".pdf"
},
{
"Name": "Pretty Good Privacy",
"Mime": "application\/pgp-encrypted",
"Ext": ""
},
{
"Name": "Pretty Good Privacy - Signature",
"Mime": "application\/pgp-signature",
"Ext": ".pgp"
},
{
"Name": "PICSRules",
"Mime": "application\/pics-rules",
"Ext": ".prf"
},
{
"Name": "PKCS #10 - Certification Request Standard",
"Mime": "application\/pkcs10",
"Ext": ".p10"
},
{
"Name": "PKCS #7 - Cryptographic Message Syntax Standard",
"Mime": "application\/pkcs7-mime",
"Ext": ".p7m"
},
{
"Name": "PKCS #7 - Cryptographic Message Syntax Standard",
"Mime": "application\/pkcs7-signature",
"Ext": ".p7s"
},
{
"Name": "PKCS #8 - Private-Key Information Syntax Standard",
"Mime": "application\/pkcs8",
"Ext": ".p8"
},
{
"Name": "Attribute Certificate",
"Mime": "application\/pkix-attr-cert",
"Ext": ".ac"
},
{
"Name": "Internet Public Key Infrastructure - Certificate",
"Mime": "application\/pkix-cert",
"Ext": ".cer"
},
{
"Name": "Internet Public Key Infrastructure - Certificate Revocation Lists",
"Mime": "application\/pkix-crl",
"Ext": ".crl"
},
{
"Name": "Internet Public Key Infrastructure - Certification Path",
"Mime": "application\/pkix-pkipath",
"Ext": ".pkipath"
},
{
"Name": "Internet Public Key Infrastructure - Certificate Management Protocole",
"Mime": "application\/pkixcmp",
"Ext": ".pki"
},
{
"Name": "Pronunciation Lexicon Specification",
"Mime": "application\/pls+xml",
"Ext": ".pls"
},
{
"Name": "PostScript",
"Mime": "application\/postscript",
"Ext": ".ai"
},
{
"Name": "CU-Writer",
"Mime": "application\/prs.cww",
"Ext": ".cww"
},
{
"Name": "Portable Symmetric Key Container",
"Mime": "application\/pskc+xml",
"Ext": ".pskcxml"
},
{
"Name": "Resource Description Framework",
"Mime": "application\/rdf+xml",
"Ext": ".rdf"
},
{
"Name": "IMS Networks",
"Mime": "application\/reginfo+xml",
"Ext": ".rif"
},
{
"Name": "Relax NG Compact Syntax",
"Mime": "application\/relax-ng-compact-syntax",
"Ext": ".rnc"
},
{
"Name": "XML Resource Lists",
"Mime": "application\/resource-lists+xml",
"Ext": ".rl"
},
{
"Name": "XML Resource Lists Diff",
"Mime": "application\/resource-lists-diff+xml",
"Ext": ".rld"
},
{
"Name": "XML Resource Lists",
"Mime": "application\/rls-services+xml",
"Ext": ".rs"
},
{
"Name": "Really Simple Discovery",
"Mime": "application\/rsd+xml",
"Ext": ".rsd"
},
{
"Name": "RSS - Really Simple Syndication",
"Mime": "application\/rss+xml",
"Ext": ".rss, .xml"
},
{
"Name": "Rich Text Format",
"Mime": "application\/rtf",
"Ext": ".rtf"
},
{
"Name": "Systems Biology Markup Language",
"Mime": "application\/sbml+xml",
"Ext": ".sbml"
},
{
"Name": "Server-Based Certificate Validation Protocol - Validation Request",
"Mime": "application\/scvp-cv-request",
"Ext": ".scq"
},
{
"Name": "Server-Based Certificate Validation Protocol - Validation Response",
"Mime": "application\/scvp-cv-response",
"Ext": ".scs"
},
{
"Name": "Server-Based Certificate Validation Protocol - Validation Policies - Request",
"Mime": "application\/scvp-vp-request",
"Ext": ".spq"
},
{
"Name": "Server-Based Certificate Validation Protocol - Validation Policies - Response",
"Mime": "application\/scvp-vp-response",
"Ext": ".spp"
},
{
"Name": "Session Description Protocol",
"Mime": "application\/sdp",
"Ext": ".sdp"
},
{
"Name": "Secure Electronic Transaction - Payment",
"Mime": "application\/set-payment-initiation",
"Ext": ".setpay"
},
{
"Name": "Secure Electronic Transaction - Registration",
"Mime": "application\/set-registration-initiation",
"Ext": ".setreg"
},
{
"Name": "S Hexdump Format",
"Mime": "application\/shf+xml",
"Ext": ".shf"
},
{
"Name": "Synchronized Multimedia Integration Language",
"Mime": "application\/smil+xml",
"Ext": ".smi"
},
{
"Name": "SPARQL - Query",
"Mime": "application\/sparql-query",
"Ext": ".rq"
},
{
"Name": "SPARQL - Results",
"Mime": "application\/sparql-results+xml",
"Ext": ".srx"
},
{
"Name": "Speech Recognition Grammar Specification",
"Mime": "application\/srgs",
"Ext": ".gram"
},
{
"Name": "Speech Recognition Grammar Specification - XML",
"Mime": "application\/srgs+xml",
"Ext": ".grxml"
},
{
"Name": "Search\/Retrieve via URL Response Format",
"Mime": "application\/sru+xml",
"Ext": ".sru"
},
{
"Name": "Speech Synthesis Markup Language",
"Mime": "application\/ssml+xml",
"Ext": ".ssml"
},
{
"Name": "Text Encoding and Interchange",
"Mime": "application\/tei+xml",
"Ext": ".tei"
},
{
"Name": "Sharing Transaction Fraud Data",
"Mime": "application\/thraud+xml",
"Ext": ".tfi"
},
{
"Name": "Time Stamped Data Envelope",
"Mime": "application\/timestamped-data",
"Ext": ".tsd"
},
{
"Name": "3rd Generation Partnership Project - Pic Large",
"Mime": "application\/vnd.3gpp.pic-bw-large",
"Ext": ".plb"
},
{
"Name": "3rd Generation Partnership Project - Pic Small",
"Mime": "application\/vnd.3gpp.pic-bw-small",
"Ext": ".psb"
},
{
"Name": "3rd Generation Partnership Project - Pic Var",
"Mime": "application\/vnd.3gpp.pic-bw-var",
"Ext": ".pvb"
},
{
"Name": "3rd Generation Partnership Project - Transaction Capabilities Application Part",
"Mime": "application\/vnd.3gpp2.tcap",
"Ext": ".tcap"
},
{
"Name": "3M Post It Notes",
"Mime": "application\/vnd.3m.post-it-notes",
"Ext": ".pwn"
},
{
"Name": "Simply Accounting",
"Mime": "application\/vnd.accpac.simply.aso",
"Ext": ".aso"
},
{
"Name": "Simply Accounting - Data Import",
"Mime": "application\/vnd.accpac.simply.imp",
"Ext": ".imp"
},
{
"Name": "ACU Cobol",
"Mime": "application\/vnd.acucobol",
"Ext": ".acu"
},
{
"Name": "ACU Cobol",
"Mime": "application\/vnd.acucorp",
"Ext": ".atc"
},
{
"Name": "Adobe AIR Application",
"Mime": "application\/vnd.adobe.air-application-installer-package+zip",
"Ext": ".air"
},
{
"Name": "Adobe Flex Project",
"Mime": "application\/vnd.adobe.fxp",
"Ext": ".fxp"
},
{
"Name": "Adobe XML Data Package",
"Mime": "application\/vnd.adobe.xdp+xml",
"Ext": ".xdp"
},
{
"Name": "Adobe XML Forms Data Format",
"Mime": "application\/vnd.adobe.xfdf",
"Ext": ".xfdf"
},
{
"Name": "Ahead AIR Application",
"Mime": "application\/vnd.ahead.space",
"Ext": ".ahead"
},
{
"Name": "AirZip FileSECURE",
"Mime": "application\/vnd.airzip.filesecure.azf",
"Ext": ".azf"
},
{
"Name": "AirZip FileSECURE",
"Mime": "application\/vnd.airzip.filesecure.azs",
"Ext": ".azs"
},
{
"Name": "Amazon Kindle eBook format",
"Mime": "application\/vnd.amazon.ebook",
"Ext": ".azw"
},
{
"Name": "Active Content Compression",
"Mime": "application\/vnd.americandynamics.acc",
"Ext": ".acc"
},
{
"Name": "AmigaDE",
"Mime": "application\/vnd.amiga.ami",
"Ext": ".ami"
},
{
"Name": "Android Package Archive",
"Mime": "application\/vnd.android.package-archive",
"Ext": ".apk"
},
{
"Name": "ANSER-WEB Terminal Client - Certificate Issue",
"Mime": "application\/vnd.anser-web-certificate-issue-initiation",
"Ext": ".cii"
},
{
"Name": "ANSER-WEB Terminal Client - Web Funds Transfer",
"Mime": "application\/vnd.anser-web-funds-transfer-initiation",
"Ext": ".fti"
},
{
"Name": "Antix Game Player",
"Mime": "application\/vnd.antix.game-component",
"Ext": ".atx"
},
{
"Name": "Apple Installer Package",
"Mime": "application\/vnd.apple.installer+xml",
"Ext": ".mpkg"
},
{
"Name": "Multimedia Playlist Unicode",
"Mime": "application\/vnd.apple.mpegurl",
"Ext": ".m3u8"
},
{
"Name": "Arista Networks Software Image",
"Mime": "application\/vnd.aristanetworks.swi",
"Ext": ".swi"
},
{
"Name": "Audiograph",
"Mime": "application\/vnd.audiograph",
"Ext": ".aep"
},
{
"Name": "Blueice Research Multipass",
"Mime": "application\/vnd.blueice.multipass",
"Ext": ".mpm"
},
{
"Name": "BMI Drawing Data Interchange",
"Mime": "application\/vnd.bmi",
"Ext": ".bmi"
},
{
"Name": "BusinessObjects",
"Mime": "application\/vnd.businessobjects",
"Ext": ".rep"
},
{
"Name": "CambridgeSoft Chem Draw",
"Mime": "application\/vnd.chemdraw+xml",
"Ext": ".cdxml"
},
{
"Name": "Karaoke on Chipnuts Chipsets",
"Mime": "application\/vnd.chipnuts.karaoke-mmd",
"Ext": ".mmd"
},
{
"Name": "Interactive Geometry Software Cinderella",
"Mime": "application\/vnd.cinderella",
"Ext": ".cdy"
},
{
"Name": "Claymore Data Files",
"Mime": "application\/vnd.claymore",
"Ext": ".cla"
},
{
"Name": "RetroPlatform Player",
"Mime": "application\/vnd.cloanto.rp9",
"Ext": ".rp9"
},
{
"Name": "Clonk Game",
"Mime": "application\/vnd.clonk.c4group",
"Ext": ".c4g"
},
{
"Name": "ClueTrust CartoMobile - Config",
"Mime": "application\/vnd.cluetrust.cartomobile-config",
"Ext": ".c11amc"
},
{
"Name": "ClueTrust CartoMobile - Config Package",
"Mime": "application\/vnd.cluetrust.cartomobile-config-pkg",
"Ext": ".c11amz"
},
{
"Name": "Sixth Floor Media - CommonSpace",
"Mime": "application\/vnd.commonspace",
"Ext": ".csp"
},
{
"Name": "CIM Database",
"Mime": "application\/vnd.contact.cmsg",
"Ext": ".cdbcmsg"
},
{
"Name": "CosmoCaller",
"Mime": "application\/vnd.cosmocaller",
"Ext": ".cmc"
},
{
"Name": "CrickSoftware - Clicker",
"Mime": "application\/vnd.crick.clicker",
"Ext": ".clkx"
},
{
"Name": "CrickSoftware - Clicker - Keyboard",
"Mime": "application\/vnd.crick.clicker.keyboard",
"Ext": ".clkk"
},
{
"Name": "CrickSoftware - Clicker - Palette",
"Mime": "application\/vnd.crick.clicker.palette",
"Ext": ".clkp"
},
{
"Name": "CrickSoftware - Clicker - Template",
"Mime": "application\/vnd.crick.clicker.template",
"Ext": ".clkt"
},
{
"Name": "CrickSoftware - Clicker - Wordbank",
"Mime": "application\/vnd.crick.clicker.wordbank",
"Ext": ".clkw"
},
{
"Name": "Critical Tools - PERT Chart EXPERT",
"Mime": "application\/vnd.criticaltools.wbs+xml",
"Ext": ".wbs"
},
{
"Name": "PosML",
"Mime": "application\/vnd.ctc-posml",
"Ext": ".pml"
},
{
"Name": "Adobe PostScript Printer Description File Format",
"Mime": "application\/vnd.cups-ppd",
"Ext": ".ppd"
},
{
"Name": "CURL Applet",
"Mime": "application\/vnd.curl.car",
"Ext": ".car"
},
{
"Name": "CURL Applet",
"Mime": "application\/vnd.curl.pcurl",
"Ext": ".pcurl"
},
{
"Name": "RemoteDocs R-Viewer",
"Mime": "application\/vnd.data-vision.rdz",
"Ext": ".rdz"
},
{
"Name": "FCS Express Layout Link",
"Mime": "application\/vnd.denovo.fcselayout-link",
"Ext": ".fe_launch"
},
{
"Name": "New Moon Liftoff\/DNA",
"Mime": "application\/vnd.dna",
"Ext": ".dna"
},
{
"Name": "Dolby Meridian Lossless Packing",
"Mime": "application\/vnd.dolby.mlp",
"Ext": ".mlp"
},
{
"Name": "DPGraph",
"Mime": "application\/vnd.dpgraph",
"Ext": ".dpg"
},
{
"Name": "DreamFactory",
"Mime": "application\/vnd.dreamfactory",
"Ext": ".dfac"
},
{
"Name": "Digital Video Broadcasting",
"Mime": "application\/vnd.dvb.ait",
"Ext": ".ait"
},
{
"Name": "Digital Video Broadcasting",
"Mime": "application\/vnd.dvb.service",
"Ext": ".svc"
},
{
"Name": "DynaGeo",
"Mime": "application\/vnd.dynageo",
"Ext": ".geo"
},
{
"Name": "EcoWin Chart",
"Mime": "application\/vnd.ecowin.chart",
"Ext": ".mag"
},
{
"Name": "Enliven Viewer",
"Mime": "application\/vnd.enliven",
"Ext": ".nml"
},
{
"Name": "QUASS Stream Player",
"Mime": "application\/vnd.epson.esf",
"Ext": ".esf"
},
{
"Name": "QUASS Stream Player",
"Mime": "application\/vnd.epson.msf",
"Ext": ".msf"
},
{
"Name": "QuickAnime Player",
"Mime": "application\/vnd.epson.quickanime",
"Ext": ".qam"
},
{
"Name": "SimpleAnimeLite Player",
"Mime": "application\/vnd.epson.salt",
"Ext": ".slt"
},
{
"Name": "QUASS Stream Player",
"Mime": "application\/vnd.epson.ssf",
"Ext": ".ssf"
},
{
"Name": "MICROSEC e-Szign\u00a2",
"Mime": "application\/vnd.eszigno3+xml",
"Ext": ".es3"
},
{
"Name": "EZPix Secure Photo Album",
"Mime": "application\/vnd.ezpix-album",
"Ext": ".ez2"
},
{
"Name": "EZPix Secure Photo Album",
"Mime": "application\/vnd.ezpix-package",
"Ext": ".ez3"
},
{
"Name": "Forms Data Format",
"Mime": "application\/vnd.fdf",
"Ext": ".fdf"
},
{
"Name": "Digital Siesmograph Networks - SEED Datafiles",
"Mime": "application\/vnd.fdsn.seed",
"Ext": ".seed"
},
{
"Name": "NpGraphIt",
"Mime": "application\/vnd.flographit",
"Ext": ".gph"
},
{
"Name": "FluxTime Clip",
"Mime": "application\/vnd.fluxtime.clip",
"Ext": ".ftc"
},
{
"Name": "FrameMaker Normal Format",
"Mime": "application\/vnd.framemaker",
"Ext": ".fm"
},
{
"Name": "Frogans Player",
"Mime": "application\/vnd.frogans.fnc",
"Ext": ".fnc"
},
{
"Name": "Frogans Player",
"Mime": "application\/vnd.frogans.ltf",
"Ext": ".ltf"
},
{
"Name": "Friendly Software Corporation",
"Mime": "application\/vnd.fsc.weblaunch",
"Ext": ".fsc"
},
{
"Name": "Fujitsu Oasys",
"Mime": "application\/vnd.fujitsu.oasys",
"Ext": ".oas"
},
{
"Name": "Fujitsu Oasys",
"Mime": "application\/vnd.fujitsu.oasys2",
"Ext": ".oa2"
},
{
"Name": "Fujitsu Oasys",
"Mime": "application\/vnd.fujitsu.oasys3",
"Ext": ".oa3"
},
{
"Name": "Fujitsu Oasys",
"Mime": "application\/vnd.fujitsu.oasysgp",
"Ext": ".fg5"
},
{
"Name": "Fujitsu Oasys",
"Mime": "application\/vnd.fujitsu.oasysprs",
"Ext": ".bh2"
},
{
"Name": "Fujitsu - Xerox 2D CAD Data",
"Mime": "application\/vnd.fujixerox.ddd",
"Ext": ".ddd"
},
{
"Name": "Fujitsu - Xerox DocuWorks",
"Mime": "application\/vnd.fujixerox.docuworks",
"Ext": ".xdw"
},
{
"Name": "Fujitsu - Xerox DocuWorks Binder",
"Mime": "application\/vnd.fujixerox.docuworks.binder",
"Ext": ".xbd"
},
{
"Name": "FuzzySheet",
"Mime": "application\/vnd.fuzzysheet",
"Ext": ".fzs"
},
{
"Name": "Genomatix Tuxedo Framework",
"Mime": "application\/vnd.genomatix.tuxedo",
"Ext": ".txd"
},
{
"Name": "GeoGebra",
"Mime": "application\/vnd.geogebra.file",
"Ext": ".ggb"
},
{
"Name": "GeoGebra",
"Mime": "application\/vnd.geogebra.tool",
"Ext": ".ggt"
},
{
"Name": "GeoMetry Explorer",
"Mime": "application\/vnd.geometry-explorer",
"Ext": ".gex"
},
{
"Name": "GEONExT and JSXGraph",
"Mime": "application\/vnd.geonext",
"Ext": ".gxt"
},
{
"Name": "GeoplanW",
"Mime": "application\/vnd.geoplan",
"Ext": ".g2w"
},
{
"Name": "GeospacW",
"Mime": "application\/vnd.geospace",
"Ext": ".g3w"
},
{
"Name": "GameMaker ActiveX",
"Mime": "application\/vnd.gmx",
"Ext": ".gmx"
},
{
"Name": "Google Earth - KML",
"Mime": "application\/vnd.google-earth.kml+xml",
"Ext": ".kml"
},
{
"Name": "Google Earth - Zipped KML",
"Mime": "application\/vnd.google-earth.kmz",
"Ext": ".kmz"
},
{
"Name": "GrafEq",
"Mime": "application\/vnd.grafeq",
"Ext": ".gqf"
},
{
"Name": "Groove - Account",
"Mime": "application\/vnd.groove-account",
"Ext": ".gac"
},
{
"Name": "Groove - Help",
"Mime": "application\/vnd.groove-help",
"Ext": ".ghf"
},
{
"Name": "Groove - Identity Message",
"Mime": "application\/vnd.groove-identity-message",
"Ext": ".gim"
},
{
"Name": "Groove - Injector",
"Mime": "application\/vnd.groove-injector",
"Ext": ".grv"
},
{
"Name": "Groove - Tool Message",
"Mime": "application\/vnd.groove-tool-message",
"Ext": ".gtm"
},
{
"Name": "Groove - Tool Template",
"Mime": "application\/vnd.groove-tool-template",
"Ext": ".tpl"
},
{
"Name": "Groove - Vcard",
"Mime": "application\/vnd.groove-vcard",
"Ext": ".vcg"
},
{
"Name": "Hypertext Application Language",
"Mime": "application\/vnd.hal+xml",
"Ext": ".hal"
},
{
"Name": "ZVUE Media Manager",
"Mime": "application\/vnd.handheld-entertainment+xml",
"Ext": ".zmm"
},
{
"Name": "Homebanking Computer Interface (HBCI)",
"Mime": "application\/vnd.hbci",
"Ext": ".hbci"
},
{
"Name": "Archipelago Lesson Player",
"Mime": "application\/vnd.hhe.lesson-player",
"Ext": ".les"
},
{
"Name": "HP-GL\/2 and HP RTL",
"Mime": "application\/vnd.hp-hpgl",
"Ext": ".hpgl"
},
{
"Name": "Hewlett Packard Instant Delivery",
"Mime": "application\/vnd.hp-hpid",
"Ext": ".hpid"
},
{
"Name": "Hewlett-Packards WebPrintSmart",
"Mime": "application\/vnd.hp-hps",
"Ext": ".hps"
},
{
"Name": "HP Indigo Digital Press - Job Layout Languate",
"Mime": "application\/vnd.hp-jlyt",
"Ext": ".jlt"
},
{
"Name": "HP Printer Command Language",
"Mime": "application\/vnd.hp-pcl",
"Ext": ".pcl"
},
{
"Name": "PCL 6 Enhanced (Formely PCL XL)",
"Mime": "application\/vnd.hp-pclxl",
"Ext": ".pclxl"
},
{
"Name": "Hydrostatix Master Suite",
"Mime": "application\/vnd.hydrostatix.sof-data",
"Ext": ".sfd-hdstx"
},
{
"Name": "3D Crossword Plugin",
"Mime": "application\/vnd.hzn-3d-crossword",
"Ext": ".x3d"
},
{
"Name": "MiniPay",
"Mime": "application\/vnd.ibm.minipay",
"Ext": ".mpy"
},
{
"Name": "MO:DCA-P",
"Mime": "application\/vnd.ibm.modcap",
"Ext": ".afp"
},
{
"Name": "IBM DB2 Rights Manager",
"Mime": "application\/vnd.ibm.rights-management",
"Ext": ".irm"
},
{
"Name": "IBM Electronic Media Management System - Secure Container",
"Mime": "application\/vnd.ibm.secure-container",
"Ext": ".sc"
},
{
"Name": "ICC profile",
"Mime": "application\/vnd.iccprofile",
"Ext": ".icc"
},
{
"Name": "igLoader",
"Mime": "application\/vnd.igloader",
"Ext": ".igl"
},
{
"Name": "ImmerVision PURE Players",
"Mime": "application\/vnd.immervision-ivp",
"Ext": ".ivp"
},
{
"Name": "ImmerVision PURE Players",
"Mime": "application\/vnd.immervision-ivu",
"Ext": ".ivu"
},
{
"Name": "IOCOM Visimeet",
"Mime": "application\/vnd.insors.igm",
"Ext": ".igm"
},
{
"Name": "Intercon FormNet",
"Mime": "application\/vnd.intercon.formnet",
"Ext": ".xpw"
},
{
"Name": "Interactive Geometry Software",
"Mime": "application\/vnd.intergeo",
"Ext": ".i2g"
},
{
"Name": "Open Financial Exchange",
"Mime": "application\/vnd.intu.qbo",
"Ext": ".qbo"
},
{
"Name": "Quicken",
"Mime": "application\/vnd.intu.qfx",
"Ext": ".qfx"
},
{
"Name": "IP Unplugged Roaming Client",
"Mime": "application\/vnd.ipunplugged.rcprofile",
"Ext": ".rcprofile"
},
{
"Name": "iRepository \/ Lucidoc Editor",
"Mime": "application\/vnd.irepository.package+xml",
"Ext": ".irp"
},
{
"Name": "Express by Infoseek",
"Mime": "application\/vnd.is-xpr",
"Ext": ".xpr"
},
{
"Name": "International Society for Advancement of Cytometry",
"Mime": "application\/vnd.isac.fcs",
"Ext": ".fcs"
},
{
"Name": "Lightspeed Audio Lab",
"Mime": "application\/vnd.jam",
"Ext": ".jam"
},
{
"Name": "Mobile Information Device Profile",
"Mime": "application\/vnd.jcp.javame.midlet-rms",
"Ext": ".rms"
},
{
"Name": "RhymBox",
"Mime": "application\/vnd.jisp",
"Ext": ".jisp"
},
{
"Name": "Joda Archive",
"Mime": "application\/vnd.joost.joda-archive",
"Ext": ".joda"
},
{
"Name": "Kahootz",
"Mime": "application\/vnd.kahootz",
"Ext": ".ktz"
},
{
"Name": "KDE KOffice Office Suite - Karbon",
"Mime": "application\/vnd.kde.karbon",
"Ext": ".karbon"
},
{
"Name": "KDE KOffice Office Suite - KChart",
"Mime": "application\/vnd.kde.kchart",
"Ext": ".chrt"
},
{
"Name": "KDE KOffice Office Suite - Kformula",
"Mime": "application\/vnd.kde.kformula",
"Ext": ".kfo"
},
{
"Name": "KDE KOffice Office Suite - Kivio",
"Mime": "application\/vnd.kde.kivio",
"Ext": ".flw"
},
{
"Name": "KDE KOffice Office Suite - Kontour",
"Mime": "application\/vnd.kde.kontour",
"Ext": ".kon"
},
{
"Name": "KDE KOffice Office Suite - Kpresenter",
"Mime": "application\/vnd.kde.kpresenter",
"Ext": ".kpr"
},
{
"Name": "KDE KOffice Office Suite - Kspread",
"Mime": "application\/vnd.kde.kspread",
"Ext": ".ksp"
},
{
"Name": "KDE KOffice Office Suite - Kword",
"Mime": "application\/vnd.kde.kword",
"Ext": ".kwd"
},
{
"Name": "Kenamea App",
"Mime": "application\/vnd.kenameaapp",
"Ext": ".htke"
},
{
"Name": "Kidspiration",
"Mime": "application\/vnd.kidspiration",
"Ext": ".kia"
},
{
"Name": "Kinar Applications",
"Mime": "application\/vnd.kinar",
"Ext": ".kne"
},
{
"Name": "SSEYO Koan Play File",
"Mime": "application\/vnd.koan",
"Ext": ".skp"
},
{
"Name": "Kodak Storyshare",
"Mime": "application\/vnd.kodak-descriptor",
"Ext": ".sse"
},
{
"Name": "Laser App Enterprise",
"Mime": "application\/vnd.las.las+xml",
"Ext": ".lasxml"
},
{
"Name": "Life Balance - Desktop Edition",
"Mime": "application\/vnd.llamagraphics.life-balance.desktop",
"Ext": ".lbd"
},
{
"Name": "Life Balance - Exchange Format",
"Mime": "application\/vnd.llamagraphics.life-balance.exchange+xml",
"Ext": ".lbe"
},
{
"Name": "Lotus 1-2-3",
"Mime": "application\/vnd.lotus-1-2-3",
"Ext": ".123"
},
{
"Name": "Lotus Approach",
"Mime": "application\/vnd.lotus-approach",
"Ext": ".apr"
},
{
"Name": "Lotus Freelance",
"Mime": "application\/vnd.lotus-freelance",
"Ext": ".pre"
},
{
"Name": "Lotus Notes",
"Mime": "application\/vnd.lotus-notes",
"Ext": ".nsf"
},
{
"Name": "Lotus Organizer",
"Mime": "application\/vnd.lotus-organizer",
"Ext": ".org"
},
{
"Name": "Lotus Screencam",
"Mime": "application\/vnd.lotus-screencam",
"Ext": ".scm"
},
{
"Name": "Lotus Wordpro",
"Mime": "application\/vnd.lotus-wordpro",
"Ext": ".lwp"
},
{
"Name": "MacPorts Port System",
"Mime": "application\/vnd.macports.portpkg",
"Ext": ".portpkg"
},
{
"Name": "Micro CADAM Helix D&D",
"Mime": "application\/vnd.mcd",
"Ext": ".mcd"
},
{
"Name": "MedCalc",
"Mime": "application\/vnd.medcalcdata",
"Ext": ".mc1"
},
{
"Name": "MediaRemote",
"Mime": "application\/vnd.mediastation.cdkey",
"Ext": ".cdkey"
},
{
"Name": "Medical Waveform Encoding Format",
"Mime": "application\/vnd.mfer",
"Ext": ".mwf"
},
{
"Name": "Melody Format for Mobile Platform",
"Mime": "application\/vnd.mfmp",
"Ext": ".mfm"
},
{
"Name": "Micrografx",
"Mime": "application\/vnd.micrografx.flo",
"Ext": ".flo"
},
{
"Name": "Micrografx iGrafx Professional",
"Mime": "application\/vnd.micrografx.igx",
"Ext": ".igx"
},
{
"Name": "FrameMaker Interchange Format",
"Mime": "application\/vnd.mif",
"Ext": ".mif"
},
{
"Name": "Mobius Management Systems - UniversalArchive",
"Mime": "application\/vnd.mobius.daf",
"Ext": ".daf"
},
{
"Name": "Mobius Management Systems - Distribution Database",
"Mime": "application\/vnd.mobius.dis",
"Ext": ".dis"
},
{
"Name": "Mobius Management Systems - Basket file",
"Mime": "application\/vnd.mobius.mbk",
"Ext": ".mbk"
},
{
"Name": "Mobius Management Systems - Query File",
"Mime": "application\/vnd.mobius.mqy",
"Ext": ".mqy"
},
{
"Name": "Mobius Management Systems - Script Language",
"Mime": "application\/vnd.mobius.msl",
"Ext": ".msl"
},
{
"Name": "Mobius Management Systems - Policy Definition Language File",
"Mime": "application\/vnd.mobius.plc",
"Ext": ".plc"
},
{
"Name": "Mobius Management Systems - Topic Index File",
"Mime": "application\/vnd.mobius.txf",
"Ext": ".txf"
},
{
"Name": "Mophun VM",
"Mime": "application\/vnd.mophun.application",
"Ext": ".mpn"
},
{
"Name": "Mophun Certificate",
"Mime": "application\/vnd.mophun.certificate",
"Ext": ".mpc"
},
{
"Name": "XUL - XML User Interface Language",
"Mime": "application\/vnd.mozilla.xul+xml",
"Ext": ".xul"
},
{
"Name": "Microsoft Artgalry",
"Mime": "application\/vnd.ms-artgalry",
"Ext": ".cil"
},
{
"Name": "Microsoft Cabinet File",
"Mime": "application\/vnd.ms-cab-compressed",
"Ext": ".cab"
},
{
"Name": "Microsoft Excel",
"Mime": "application\/vnd.ms-excel",
"Ext": ".xls"
},
{
"Name": "Microsoft Excel - Add-In File",
"Mime": "application\/vnd.ms-excel.addin.macroenabled.12",
"Ext": ".xlam"
},
{
"Name": "Microsoft Excel - Binary Workbook",
"Mime": "application\/vnd.ms-excel.sheet.binary.macroenabled.12",
"Ext": ".xlsb"
},
{
"Name": "Microsoft Excel - Macro-Enabled Workbook",
"Mime": "application\/vnd.ms-excel.sheet.macroenabled.12",
"Ext": ".xlsm"
},
{
"Name": "Microsoft Excel - Macro-Enabled Template File",
"Mime": "application\/vnd.ms-excel.template.macroenabled.12",
"Ext": ".xltm"
},
{
"Name": "Microsoft Embedded OpenType",
"Mime": "application\/vnd.ms-fontobject",
"Ext": ".eot"
},
{
"Name": "Microsoft Html Help File",
"Mime": "application\/vnd.ms-htmlhelp",
"Ext": ".chm"
},
{
"Name": "Microsoft Class Server",
"Mime": "application\/vnd.ms-ims",
"Ext": ".ims"
},
{
"Name": "Microsoft Learning Resource Module",
"Mime": "application\/vnd.ms-lrm",
"Ext": ".lrm"
},
{
"Name": "Microsoft Office System Release Theme",
"Mime": "application\/vnd.ms-officetheme",
"Ext": ".thmx"
},
{
"Name": "Microsoft Trust UI Provider - Security Catalog",
"Mime": "application\/vnd.ms-pki.seccat",
"Ext": ".cat"
},
{
"Name": "Microsoft Trust UI Provider - Certificate Trust Link",
"Mime": "application\/vnd.ms-pki.stl",
"Ext": ".stl"
},
{
"Name": "Microsoft PowerPoint",
"Mime": "application\/vnd.ms-powerpoint",
"Ext": ".ppt"
},
{
"Name": "Microsoft PowerPoint - Add-in file",
"Mime": "application\/vnd.ms-powerpoint.addin.macroenabled.12",
"Ext": ".ppam"
},
{
"Name": "Microsoft PowerPoint - Macro-Enabled Presentation File",
"Mime": "application\/vnd.ms-powerpoint.presentation.macroenabled.12",
"Ext": ".pptm"
},
{
"Name": "Microsoft PowerPoint - Macro-Enabled Open XML Slide",
"Mime": "application\/vnd.ms-powerpoint.slide.macroenabled.12",
"Ext": ".sldm"
},
{
"Name": "Microsoft PowerPoint - Macro-Enabled Slide Show File",
"Mime": "application\/vnd.ms-powerpoint.slideshow.macroenabled.12",
"Ext": ".ppsm"
},
{
"Name": "Micosoft PowerPoint - Macro-Enabled Template File",
"Mime": "application\/vnd.ms-powerpoint.template.macroenabled.12",
"Ext": ".potm"
},
{
"Name": "Microsoft Project",
"Mime": "application\/vnd.ms-project",
"Ext": ".mpp"
},
{
"Name": "Micosoft Word - Macro-Enabled Document",
"Mime": "application\/vnd.ms-word.document.macroenabled.12",
"Ext": ".docm"
},
{
"Name": "Micosoft Word - Macro-Enabled Template",
"Mime": "application\/vnd.ms-word.template.macroenabled.12",
"Ext": ".dotm"
},
{
"Name": "Microsoft Works",
"Mime": "application\/vnd.ms-works",
"Ext": ".wps"
},
{
"Name": "Microsoft Windows Media Player Playlist",
"Mime": "application\/vnd.ms-wpl",
"Ext": ".wpl"
},
{
"Name": "Microsoft XML Paper Specification",
"Mime": "application\/vnd.ms-xpsdocument",
"Ext": ".xps"
},
{
"Name": "3GPP MSEQ File",
"Mime": "application\/vnd.mseq",
"Ext": ".mseq"
},
{
"Name": "MUsical Score Interpreted Code Invented  for the ASCII designation of Notation",
"Mime": "application\/vnd.musician",
"Ext": ".mus"
},
{
"Name": "Muvee Automatic Video Editing",
"Mime": "application\/vnd.muvee.style",
"Ext": ".msty"
},
{
"Name": "neuroLanguage",
"Mime": "application\/vnd.neurolanguage.nlu",
"Ext": ".nlu"
},
{
"Name": "NobleNet Directory",
"Mime": "application\/vnd.noblenet-directory",
"Ext": ".nnd"
},
{
"Name": "NobleNet Sealer",
"Mime": "application\/vnd.noblenet-sealer",
"Ext": ".nns"
},
{
"Name": "NobleNet Web",
"Mime": "application\/vnd.noblenet-web",
"Ext": ".nnw"
},
{
"Name": "N-Gage Game Data",
"Mime": "application\/vnd.nokia.n-gage.data",
"Ext": ".ngdat"
},
{
"Name": "N-Gage Game Installer",
"Mime": "application\/vnd.nokia.n-gage.symbian.install",
"Ext": ".n-gage"
},
{
"Name": "Nokia Radio Application - Preset",
"Mime": "application\/vnd.nokia.radio-preset",
"Ext": ".rpst"
},
{
"Name": "Nokia Radio Application - Preset",
"Mime": "application\/vnd.nokia.radio-presets",
"Ext": ".rpss"
},
{
"Name": "Novadigms RADIA and EDM products",
"Mime": "application\/vnd.novadigm.edm",
"Ext": ".edm"
},
{
"Name": "Novadigms RADIA and EDM products",
"Mime": "application\/vnd.novadigm.edx",
"Ext": ".edx"
},
{
"Name": "Novadigms RADIA and EDM products",
"Mime": "application\/vnd.novadigm.ext",
"Ext": ".ext"
},
{
"Name": "OpenDocument Chart",
"Mime": "application\/vnd.oasis.opendocument.chart",
"Ext": ".odc"
},
{
"Name": "OpenDocument Chart Template",
"Mime": "application\/vnd.oasis.opendocument.chart-template",
"Ext": ".otc"
},
{
"Name": "OpenDocument Database",
"Mime": "application\/vnd.oasis.opendocument.database",
"Ext": ".odb"
},
{
"Name": "OpenDocument Formula",
"Mime": "application\/vnd.oasis.opendocument.formula",
"Ext": ".odf"
},
{
"Name": "OpenDocument Formula Template",
"Mime": "application\/vnd.oasis.opendocument.formula-template",
"Ext": ".odft"
},
{
"Name": "OpenDocument Graphics",
"Mime": "application\/vnd.oasis.opendocument.graphics",
"Ext": ".odg"
},
{
"Name": "OpenDocument Graphics Template",
"Mime": "application\/vnd.oasis.opendocument.graphics-template",
"Ext": ".otg"
},
{
"Name": "OpenDocument Image",
"Mime": "application\/vnd.oasis.opendocument.image",
"Ext": ".odi"
},
{
"Name": "OpenDocument Image Template",
"Mime": "application\/vnd.oasis.opendocument.image-template",
"Ext": ".oti"
},
{
"Name": "OpenDocument Presentation",
"Mime": "application\/vnd.oasis.opendocument.presentation",
"Ext": ".odp"
},
{
"Name": "OpenDocument Presentation Template",
"Mime": "application\/vnd.oasis.opendocument.presentation-template",
"Ext": ".otp"
},
{
"Name": "OpenDocument Spreadsheet",
"Mime": "application\/vnd.oasis.opendocument.spreadsheet",
"Ext": ".ods"
},
{
"Name": "OpenDocument Spreadsheet Template",
"Mime": "application\/vnd.oasis.opendocument.spreadsheet-template",
"Ext": ".ots"
},
{
"Name": "OpenDocument Text",
"Mime": "application\/vnd.oasis.opendocument.text",
"Ext": ".odt"
},
{
"Name": "OpenDocument Text Master",
"Mime": "application\/vnd.oasis.opendocument.text-master",
"Ext": ".odm"
},
{
"Name": "OpenDocument Text Template",
"Mime": "application\/vnd.oasis.opendocument.text-template",
"Ext": ".ott"
},
{
"Name": "Open Document Text Web",
"Mime": "application\/vnd.oasis.opendocument.text-web",
"Ext": ".oth"
},
{
"Name": "Sugar Linux Application Bundle",
"Mime": "application\/vnd.olpc-sugar",
"Ext": ".xo"
},
{
"Name": "OMA Download Agents",
"Mime": "application\/vnd.oma.dd2+xml",
"Ext": ".dd2"
},
{
"Name": "Open Office Extension",
"Mime": "application\/vnd.openofficeorg.extension",
"Ext": ".oxt"
},
{
"Name": "Microsoft Office - OOXML - Presentation",
"Mime": "application\/vnd.openxmlformats-officedocument.presentationml.presentation",
"Ext": ".pptx"
},
{
"Name": "Microsoft Office - OOXML - Presentation (Slide)",
"Mime": "application\/vnd.openxmlformats-officedocument.presentationml.slide",
"Ext": ".sldx"
},
{
"Name": "Microsoft Office - OOXML - Presentation (Slideshow)",
"Mime": "application\/vnd.openxmlformats-officedocument.presentationml.slideshow",
"Ext": ".ppsx"
},
{
"Name": "Microsoft Office - OOXML - Presentation Template",
"Mime": "application\/vnd.openxmlformats-officedocument.presentationml.template",
"Ext": ".potx"
},
{
"Name": "Microsoft Office - OOXML - Spreadsheet",
"Mime": "application\/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
"Ext": ".xlsx"
},
{
"Name": "Microsoft Office - OOXML - Spreadsheet Teplate",
"Mime": "application\/vnd.openxmlformats-officedocument.spreadsheetml.template",
"Ext": ".xltx"
},
{
"Name": "Microsoft Office - OOXML - Word Document",
"Mime": "application\/vnd.openxmlformats-officedocument.wordprocessingml.document",
"Ext": ".docx"
},
{
"Name": "Microsoft Office - OOXML - Word Document Template",
"Mime": "application\/vnd.openxmlformats-officedocument.wordprocessingml.template",
"Ext": ".dotx"
},
{
"Name": "MapGuide DBXML",
"Mime": "application\/vnd.osgeo.mapguide.package",
"Ext": ".mgp"
},
{
"Name": "OSGi Deployment Package",
"Mime": "application\/vnd.osgi.dp",
"Ext": ".dp"
},
{
"Name": "PalmOS Data",
"Mime": "application\/vnd.palm",
"Ext": ".pdb"
},
{
"Name": "PawaaFILE",
"Mime": "application\/vnd.pawaafile",
"Ext": ".paw"
},
{
"Name": "Proprietary P&G Standard Reporting System",
"Mime": "application\/vnd.pg.format",
"Ext": ".str"
},
{
"Name": "Proprietary P&G Standard Reporting System",
"Mime": "application\/vnd.pg.osasli",
"Ext": ".ei6"
},
{
"Name": "Pcsel eFIF File",
"Mime": "application\/vnd.picsel",
"Ext": ".efif"
},
{
"Name": "Qualcomms Plaza Mobile Internet",
"Mime": "application\/vnd.pmi.widget",
"Ext": ".wg"
},
{
"Name": "PocketLearn Viewers",
"Mime": "application\/vnd.pocketlearn",
"Ext": ".plf"
},
{
"Name": "PowerBuilder",
"Mime": "application\/vnd.powerbuilder6",
"Ext": ".pbd"
},
{
"Name": "Preview Systems ZipLock\/VBox",
"Mime": "application\/vnd.previewsystems.box",
"Ext": ".box"
},
{
"Name": "EFI Proteus",
"Mime": "application\/vnd.proteus.magazine",
"Ext": ".mgz"
},
{
"Name": "PubliShare Objects",
"Mime": "application\/vnd.publishare-delta-tree",
"Ext": ".qps"
},
{
"Name": "Princeton Video Image",
"Mime": "application\/vnd.pvi.ptid1",
"Ext": ".ptid"
},
{
"Name": "QuarkXpress",
"Mime": "application\/vnd.quark.quarkxpress",
"Ext": ".qxd"
},
{
"Name": "RealVNC",
"Mime": "application\/vnd.realvnc.bed",
"Ext": ".bed"
},
{
"Name": "Recordare Applications",
"Mime": "application\/vnd.recordare.musicxml",
"Ext": ".mxl"
},
{
"Name": "Recordare Applications",
"Mime": "application\/vnd.recordare.musicxml+xml",
"Ext": ".musicxml"
},
{
"Name": "CryptoNote",
"Mime": "application\/vnd.rig.cryptonote",
"Ext": ".cryptonote"
},
{
"Name": "Blackberry COD File",
"Mime": "application\/vnd.rim.cod",
"Ext": ".cod"
},
{
"Name": "RealMedia",
"Mime": "application\/vnd.rn-realmedia",
"Ext": ".rm"
},
{
"Name": "ROUTE 66 Location Based Services",
"Mime": "application\/vnd.route66.link66+xml",
"Ext": ".link66"
},
{
"Name": "SailingTracker",
"Mime": "application\/vnd.sailingtracker.track",
"Ext": ".st"
},
{
"Name": "SeeMail",
"Mime": "application\/vnd.seemail",
"Ext": ".see"
},
{
"Name": "Secured eMail",
"Mime": "application\/vnd.sema",
"Ext": ".sema"
},
{
"Name": "Secured eMail",
"Mime": "application\/vnd.semd",
"Ext": ".semd"
},
{
"Name": "Secured eMail",
"Mime": "application\/vnd.semf",
"Ext": ".semf"
},
{
"Name": "Shana Informed Filler",
"Mime": "application\/vnd.shana.informed.formdata",
"Ext": ".ifm"
},
{
"Name": "Shana Informed Filler",
"Mime": "application\/vnd.shana.informed.formtemplate",
"Ext": ".itp"
},
{
"Name": "Shana Informed Filler",
"Mime": "application\/vnd.shana.informed.interchange",
"Ext": ".iif"
},
{
"Name": "Shana Informed Filler",
"Mime": "application\/vnd.shana.informed.package",
"Ext": ".ipk"
},
{
"Name": "SimTech MindMapper",
"Mime": "application\/vnd.simtech-mindmapper",
"Ext": ".twd"
},
{
"Name": "SMAF File",
"Mime": "application\/vnd.smaf",
"Ext": ".mmf"
},
{
"Name": "SMART Technologies Apps",
"Mime": "application\/vnd.smart.teacher",
"Ext": ".teacher"
},
{
"Name": "SudokuMagic",
"Mime": "application\/vnd.solent.sdkm+xml",
"Ext": ".sdkm"
},
{
"Name": "TIBCO Spotfire",
"Mime": "application\/vnd.spotfire.dxp",
"Ext": ".dxp"
},
{
"Name": "TIBCO Spotfire",
"Mime": "application\/vnd.spotfire.sfs",
"Ext": ".sfs"
},
{
"Name": "StarOffice - Calc",
"Mime": "application\/vnd.stardivision.calc",
"Ext": ".sdc"
},
{
"Name": "StarOffice - Draw",
"Mime": "application\/vnd.stardivision.draw",
"Ext": ".sda"
},
{
"Name": "StarOffice - Impress",
"Mime": "application\/vnd.stardivision.impress",
"Ext": ".sdd"
},
{
"Name": "StarOffice - Math",
"Mime": "application\/vnd.stardivision.math",
"Ext": ".smf"
},
{
"Name": "StarOffice - Writer",
"Mime": "application\/vnd.stardivision.writer",
"Ext": ".sdw"
},
{
"Name": "StarOffice - Writer  (Global)",
"Mime": "application\/vnd.stardivision.writer-global",
"Ext": ".sgl"
},
{
"Name": "StepMania",
"Mime": "application\/vnd.stepmania.stepchart",
"Ext": ".sm"
},
{
"Name": "OpenOffice - Calc (Spreadsheet)",
"Mime": "application\/vnd.sun.xml.calc",
"Ext": ".sxc"
},
{
"Name": "OpenOffice - Calc Template (Spreadsheet)",
"Mime": "application\/vnd.sun.xml.calc.template",
"Ext": ".stc"
},
{
"Name": "OpenOffice - Draw (Graphics)",
"Mime": "application\/vnd.sun.xml.draw",
"Ext": ".sxd"
},
{
"Name": "OpenOffice - Draw Template (Graphics)",
"Mime": "application\/vnd.sun.xml.draw.template",
"Ext": ".std"
},
{
"Name": "OpenOffice - Impress (Presentation)",
"Mime": "application\/vnd.sun.xml.impress",
"Ext": ".sxi"
},
{
"Name": "OpenOffice - Impress Template (Presentation)",
"Mime": "application\/vnd.sun.xml.impress.template",
"Ext": ".sti"
},
{
"Name": "OpenOffice - Math (Formula)",
"Mime": "application\/vnd.sun.xml.math",
"Ext": ".sxm"
},
{
"Name": "OpenOffice - Writer (Text - HTML)",
"Mime": "application\/vnd.sun.xml.writer",
"Ext": ".sxw"
},
{
"Name": "OpenOffice - Writer (Text - HTML)",
"Mime": "application\/vnd.sun.xml.writer.global",
"Ext": ".sxg"
},
{
"Name": "OpenOffice - Writer Template (Text - HTML)",
"Mime": "application\/vnd.sun.xml.writer.template",
"Ext": ".stw"
},
{
"Name": "ScheduleUs",
"Mime": "application\/vnd.sus-calendar",
"Ext": ".sus"
},
{
"Name": "SourceView Document",
"Mime": "application\/vnd.svd",
"Ext": ".svd"
},
{
"Name": "Symbian Install Package",
"Mime": "application\/vnd.symbian.install",
"Ext": ".sis"
},
{
"Name": "SyncML",
"Mime": "application\/vnd.syncml+xml",
"Ext": ".xsm"
},
{
"Name": "SyncML - Device Management",
"Mime": "application\/vnd.syncml.dm+wbxml",
"Ext": ".bdm"
},
{
"Name": "SyncML - Device Management",
"Mime": "application\/vnd.syncml.dm+xml",
"Ext": ".xdm"
},
{
"Name": "Tao Intent",
"Mime": "application\/vnd.tao.intent-module-archive",
"Ext": ".tao"
},
{
"Name": "MobileTV",
"Mime": "application\/vnd.tmobile-livetv",
"Ext": ".tmo"
},
{
"Name": "TRI Systems Config",
"Mime": "application\/vnd.trid.tpt",
"Ext": ".tpt"
},
{
"Name": "Triscape Map Explorer",
"Mime": "application\/vnd.triscape.mxs",
"Ext": ".mxs"
},
{
"Name": "True BASIC",
"Mime": "application\/vnd.trueapp",
"Ext": ".tra"
},
{
"Name": "Universal Forms Description Language",
"Mime": "application\/vnd.ufdl",
"Ext": ".ufd"
},
{
"Name": "User Interface Quartz - Theme (Symbian)",
"Mime": "application\/vnd.uiq.theme",
"Ext": ".utz"
},
{
"Name": "UMAJIN",
"Mime": "application\/vnd.umajin",
"Ext": ".umj"
},
{
"Name": "Unity 3d",
"Mime": "application\/vnd.unity",
"Ext": ".unityweb"
},
{
"Name": "Unique Object Markup Language",
"Mime": "application\/vnd.uoml+xml",
"Ext": ".uoml"
},
{
"Name": "VirtualCatalog",
"Mime": "application\/vnd.vcx",
"Ext": ".vcx"
},
{
"Name": "Microsoft Visio",
"Mime": "application\/vnd.visio",
"Ext": ".vsd"
},
{
"Name": "Visionary",
"Mime": "application\/vnd.visionary",
"Ext": ".vis"
},
{
"Name": "Viewport+",
"Mime": "application\/vnd.vsf",
"Ext": ".vsf"
},
{
"Name": "WAP Binary XML (WBXML)",
"Mime": "application\/vnd.wap.wbxml",
"Ext": ".wbxml"
},
{
"Name": "Compiled Wireless Markup Language (WMLC)",
"Mime": "application\/vnd.wap.wmlc",
"Ext": ".wmlc"
},
{
"Name": "WMLScript",
"Mime": "application\/vnd.wap.wmlscriptc",
"Ext": ".wmlsc"
},
{
"Name": "WebTurbo",
"Mime": "application\/vnd.webturbo",
"Ext": ".wtb"
},
{
"Name": "Mathematica Notebook Player",
"Mime": "application\/vnd.wolfram.player",
"Ext": ".nbp"
},
{
"Name": "Wordperfect",
"Mime": "application\/vnd.wordperfect",
"Ext": ".wpd"
},
{
"Name": "SundaHus WQ",
"Mime": "application\/vnd.wqd",
"Ext": ".wqd"
},
{
"Name": "Worldtalk",
"Mime": "application\/vnd.wt.stf",
"Ext": ".stf"
},
{
"Name": "CorelXARA",
"Mime": "application\/vnd.xara",
"Ext": ".xar"
},
{
"Name": "Extensible Forms Description Language",
"Mime": "application\/vnd.xfdl",
"Ext": ".xfdl"
},
{
"Name": "HV Voice Dictionary",
"Mime": "application\/vnd.yamaha.hv-dic",
"Ext": ".hvd"
},
{
"Name": "HV Script",
"Mime": "application\/vnd.yamaha.hv-script",
"Ext": ".hvs"
},
{
"Name": "HV Voice Parameter",
"Mime": "application\/vnd.yamaha.hv-voice",
"Ext": ".hvp"
},
{
"Name": "Open Score Format",
"Mime": "application\/vnd.yamaha.openscoreformat",
"Ext": ".osf"
},
{
"Name": "OSFPVG",
"Mime": "application\/vnd.yamaha.openscoreformat.osfpvg+xml",
"Ext": ".osfpvg"
},
{
"Name": "SMAF Audio",
"Mime": "application\/vnd.yamaha.smaf-audio",
"Ext": ".saf"
},
{
"Name": "SMAF Phrase",
"Mime": "application\/vnd.yamaha.smaf-phrase",
"Ext": ".spf"
},
{
"Name": "CustomMenu",
"Mime": "application\/vnd.yellowriver-custom-menu",
"Ext": ".cmp"
},
{
"Name": "Z.U.L. Geometry",
"Mime": "application\/vnd.zul",
"Ext": ".zir"
},
{
"Name": "Zzazz Deck",
"Mime": "application\/vnd.zzazz.deck+xml",
"Ext": ".zaz"
},
{
"Name": "VoiceXML",
"Mime": "application\/voicexml+xml",
"Ext": ".vxml"
},
{
"Name": "Widget Packaging and XML Configuration",
"Mime": "application\/widget",
"Ext": ".wgt"
},
{
"Name": "WinHelp",
"Mime": "application\/winhlp",
"Ext": ".hlp"
},
{
"Name": "WSDL - Web Services Description Language",
"Mime": "application\/wsdl+xml",
"Ext": ".wsdl"
},
{
"Name": "Web Services Policy",
"Mime": "application\/wspolicy+xml",
"Ext": ".wspolicy"
},
{
"Name": "7-Zip",
"Mime": "application\/x-7z-compressed",
"Ext": ".7z"
},
{
"Name": "AbiWord",
"Mime": "application\/x-abiword",
"Ext": ".abw"
},
{
"Name": "Ace Archive",
"Mime": "application\/x-ace-compressed",
"Ext": ".ace"
},
{
"Name": "Adobe (Macropedia) Authorware - Binary File",
"Mime": "application\/x-authorware-bin",
"Ext": ".aab"
},
{
"Name": "Adobe (Macropedia) Authorware - Map",
"Mime": "application\/x-authorware-map",
"Ext": ".aam"
},
{
"Name": "Adobe (Macropedia) Authorware - Segment File",
"Mime": "application\/x-authorware-seg",
"Ext": ".aas"
},
{
"Name": "Binary CPIO Archive",
"Mime": "application\/x-bcpio",
"Ext": ".bcpio"
},
{
"Name": "BitTorrent",
"Mime": "application\/x-bittorrent",
"Ext": ".torrent"
},
{
"Name": "Bzip Archive",
"Mime": "application\/x-bzip",
"Ext": ".bz"
},
{
"Name": "Bzip2 Archive",
"Mime": "application\/x-bzip2",
"Ext": ".bz2"
},
{
"Name": "Video CD",
"Mime": "application\/x-cdlink",
"Ext": ".vcd"
},
{
"Name": "pIRCh",
"Mime": "application\/x-chat",
"Ext": ".chat"
},
{
"Name": "Portable Game Notation (Chess Games)",
"Mime": "application\/x-chess-pgn",
"Ext": ".pgn"
},
{
"Name": "CPIO Archive",
"Mime": "application\/x-cpio",
"Ext": ".cpio"
},
{
"Name": "C Shell Script",
"Mime": "application\/x-csh",
"Ext": ".csh"
},
{
"Name": "Debian Package",
"Mime": "application\/x-debian-package",
"Ext": ".deb"
},
{
"Name": "Adobe Shockwave Player",
"Mime": "application\/x-director",
"Ext": ".dir"
},
{
"Name": "Doom Video Game",
"Mime": "application\/x-doom",
"Ext": ".wad"
},
{
"Name": "Navigation Control file for XML (for ePub)",
"Mime": "application\/x-dtbncx+xml",
"Ext": ".ncx"
},
{
"Name": "Digital Talking Book",
"Mime": "application\/x-dtbook+xml",
"Ext": ".dtb"
},
{
"Name": "Digital Talking Book - Resource File",
"Mime": "application\/x-dtbresource+xml",
"Ext": ".res"
},
{
"Name": "Device Independent File Format (DVI)",
"Mime": "application\/x-dvi",
"Ext": ".dvi"
},
{
"Name": "Glyph Bitmap Distribution Format",
"Mime": "application\/x-font-bdf",
"Ext": ".bdf"
},
{
"Name": "Ghostscript Font",
"Mime": "application\/x-font-ghostscript",
"Ext": ".gsf"
},
{
"Name": "PSF Fonts",
"Mime": "application\/x-font-linux-psf",
"Ext": ".psf"
},
{
"Name": "OpenType Font File",
"Mime": "application\/x-font-otf",
"Ext": ".otf"
},
{
"Name": "Portable Compiled Format",
"Mime": "application\/x-font-pcf",
"Ext": ".pcf"
},
{
"Name": "Server Normal Format",
"Mime": "application\/x-font-snf",
"Ext": ".snf"
},
{
"Name": "TrueType Font",
"Mime": "application\/x-font-ttf",
"Ext": ".ttf"
},
{
"Name": "PostScript Fonts",
"Mime": "application\/x-font-type1",
"Ext": ".pfa"
},
{
"Name": "Web Open Font Format",
"Mime": "application\/x-font-woff",
"Ext": ".woff"
},
{
"Name": "FutureSplash Animator",
"Mime": "application\/x-futuresplash",
"Ext": ".spl"
},
{
"Name": "Gnumeric",
"Mime": "application\/x-gnumeric",
"Ext": ".gnumeric"
},
{
"Name": "GNU Tar Files",
"Mime": "application\/x-gtar",
"Ext": ".gtar"
},
{
"Name": "Hierarchical Data Format",
"Mime": "application\/x-hdf",
"Ext": ".hdf"
},
{
"Name": "Java Network Launching Protocol",
"Mime": "application\/x-java-jnlp-file",
"Ext": ".jnlp"
},
{
"Name": "LaTeX",
"Mime": "application\/x-latex",
"Ext": ".latex"
},
{
"Name": "Mobipocket",
"Mime": "application\/x-mobipocket-ebook",
"Ext": ".prc"
},
{
"Name": "Microsoft ClickOnce",
"Mime": "application\/x-ms-application",
"Ext": ".application"
},
{
"Name": "Microsoft Windows Media Player Download Package",
"Mime": "application\/x-ms-wmd",
"Ext": ".wmd"
},
{
"Name": "Microsoft Windows Media Player Skin Package",
"Mime": "application\/x-ms-wmz",
"Ext": ".wmz"
},
{
"Name": "Microsoft XAML Browser Application",
"Mime": "application\/x-ms-xbap",
"Ext": ".xbap"
},
{
"Name": "Microsoft Access",
"Mime": "application\/x-msaccess",
"Ext": ".mdb"
},
{
"Name": "Microsoft Office Binder",
"Mime": "application\/x-msbinder",
"Ext": ".obd"
},
{
"Name": "Microsoft Information Card",
"Mime": "application\/x-mscardfile",
"Ext": ".crd"
},
{
"Name": "Microsoft Clipboard Clip",
"Mime": "application\/x-msclip",
"Ext": ".clp"
},
{
"Name": "Microsoft Application",
"Mime": "application\/x-msdownload",
"Ext": ".exe"
},
{
"Name": "Microsoft MediaView",
"Mime": "application\/x-msmediaview",
"Ext": ".mvb"
},
{
"Name": "Microsoft Windows Metafile",
"Mime": "application\/x-msmetafile",
"Ext": ".wmf"
},
{
"Name": "Microsoft Money",
"Mime": "application\/x-msmoney",
"Ext": ".mny"
},
{
"Name": "Microsoft Publisher",
"Mime": "application\/x-mspublisher",
"Ext": ".pub"
},
{
"Name": "Microsoft Schedule+",
"Mime": "application\/x-msschedule",
"Ext": ".scd"
},
{
"Name": "Microsoft Windows Terminal Services",
"Mime": "application\/x-msterminal",
"Ext": ".trm"
},
{
"Name": "Microsoft Wordpad",
"Mime": "application\/x-mswrite",
"Ext": ".wri"
},
{
"Name": "Network Common Data Form (NetCDF)",
"Mime": "application\/x-netcdf",
"Ext": ".nc"
},
{
"Name": "PKCS #12 - Personal Information Exchange Syntax Standard",
"Mime": "application\/x-pkcs12",
"Ext": ".p12"
},
{
"Name": "PKCS #7 - Cryptographic Message Syntax Standard (Certificates)",
"Mime": "application\/x-pkcs7-certificates",
"Ext": ".p7b"
},
{
"Name": "PKCS #7 - Cryptographic Message Syntax Standard (Certificate Request Response)",
"Mime": "application\/x-pkcs7-certreqresp",
"Ext": ".p7r"
},
{
"Name": "RAR Archive",
"Mime": "application\/x-rar-compressed",
"Ext": ".rar"
},
{
"Name": "Bourne Shell Script",
"Mime": "application\/x-sh",
"Ext": ".sh"
},
{
"Name": "Shell Archive",
"Mime": "application\/x-shar",
"Ext": ".shar"
},
{
"Name": "Adobe Flash",
"Mime": "application\/x-shockwave-flash",
"Ext": ".swf"
},
{
"Name": "Microsoft Silverlight",
"Mime": "application\/x-silverlight-app",
"Ext": ".xap"
},
{
"Name": "Stuffit Archive",
"Mime": "application\/x-stuffit",
"Ext": ".sit"
},
{
"Name": "Stuffit Archive",
"Mime": "application\/x-stuffitx",
"Ext": ".sitx"
},
{
"Name": "System V Release 4 CPIO Archive",
"Mime": "application\/x-sv4cpio",
"Ext": ".sv4cpio"
},
{
"Name": "System V Release 4 CPIO Checksum Data",
"Mime": "application\/x-sv4crc",
"Ext": ".sv4crc"
},
{
"Name": "Tar File (Tape Archive)",
"Mime": "application\/x-tar",
"Ext": ".tar"
},
{
"Name": "Tcl Script",
"Mime": "application\/x-tcl",
"Ext": ".tcl"
},
{
"Name": "TeX",
"Mime": "application\/x-tex",
"Ext": ".tex"
},
{
"Name": "TeX Font Metric",
"Mime": "application\/x-tex-tfm",
"Ext": ".tfm"
},
{
"Name": "GNU Texinfo Document",
"Mime": "application\/x-texinfo",
"Ext": ".texinfo"
},
{
"Name": "Ustar (Uniform Standard Tape Archive)",
"Mime": "application\/x-ustar",
"Ext": ".ustar"
},
{
"Name": "WAIS Source",
"Mime": "application\/x-wais-source",
"Ext": ".src"
},
{
"Name": "X.509 Certificate",
"Mime": "application\/x-x509-ca-cert",
"Ext": ".der"
},
{
"Name": "Xfig",
"Mime": "application\/x-xfig",
"Ext": ".fig"
},
{
"Name": "XPInstall - Mozilla",
"Mime": "application\/x-xpinstall",
"Ext": ".xpi"
},
{
"Name": "XML Configuration Access Protocol - XCAP Diff",
"Mime": "application\/xcap-diff+xml",
"Ext": ".xdf"
},
{
"Name": "XML Encryption Syntax and Processing",
"Mime": "application\/xenc+xml",
"Ext": ".xenc"
},
{
"Name": "XHTML - The Extensible HyperText Markup Language",
"Mime": "application\/xhtml+xml",
"Ext": ".xhtml"
},
{
"Name": "XML - Extensible Markup Language",
"Mime": "application\/xml",
"Ext": ".xml"
},
{
"Name": "Document Type Definition",
"Mime": "application\/xml-dtd",
"Ext": ".dtd"
},
{
"Name": "XML-Binary Optimized Packaging",
"Mime": "application\/xop+xml",
"Ext": ".xop"
},
{
"Name": "XML Transformations",
"Mime": "application\/xslt+xml",
"Ext": ".xslt"
},
{
"Name": "XSPF - XML Shareable Playlist Format",
"Mime": "application\/xspf+xml",
"Ext": ".xspf"
},
{
"Name": "MXML",
"Mime": "application\/xv+xml",
"Ext": ".mxml"
},
{
"Name": "YANG Data Modeling Language",
"Mime": "application\/yang",
"Ext": ".yang"
},
{
"Name": "YIN (YANG - XML)",
"Mime": "application\/yin+xml",
"Ext": ".yin"
},
{
"Name": "Zip Archive",
"Mime": "application\/zip",
"Ext": ".zip"
},
{
"Name": "Adaptive differential pulse-code modulation",
"Mime": "audio\/adpcm",
"Ext": ".adp"
},
{
"Name": "Sun Audio - Au file format",
"Mime": "audio\/basic",
"Ext": ".au"
},
{
"Name": "MIDI - Musical Instrument Digital Interface",
"Mime": "audio\/midi",
"Ext": ".mid"
},
{
"Name": "MPEG-4 Audio",
"Mime": "audio\/mp4",
"Ext": ".mp4a"
},
{
"Name": "MPEG Audio",
"Mime": "audio\/mpeg",
"Ext": ".mpga"
},
{
"Name": "Ogg Audio",
"Mime": "audio\/ogg",
"Ext": ".oga"
},
{
"Name": "DECE Audio",
"Mime": "audio\/vnd.dece.audio",
"Ext": ".uva"
},
{
"Name": "Digital Winds Music",
"Mime": "audio\/vnd.digital-winds",
"Ext": ".eol"
},
{
"Name": "DRA Audio",
"Mime": "audio\/vnd.dra",
"Ext": ".dra"
},
{
"Name": "DTS Audio",
"Mime": "audio\/vnd.dts",
"Ext": ".dts"
},
{
"Name": "DTS High Definition Audio",
"Mime": "audio\/vnd.dts.hd",
"Ext": ".dtshd"
},
{
"Name": "Lucent Voice",
"Mime": "audio\/vnd.lucent.voice",
"Ext": ".lvp"
},
{
"Name": "Microsoft PlayReady Ecosystem",
"Mime": "audio\/vnd.ms-playready.media.pya",
"Ext": ".pya"
},
{
"Name": "Nuera ECELP 4800",
"Mime": "audio\/vnd.nuera.ecelp4800",
"Ext": ".ecelp4800"
},
{
"Name": "Nuera ECELP 7470",
"Mime": "audio\/vnd.nuera.ecelp7470",
"Ext": ".ecelp7470"
},
{
"Name": "Nuera ECELP 9600",
"Mime": "audio\/vnd.nuera.ecelp9600",
"Ext": ".ecelp9600"
},
{
"Name": "Hit n Mix",
"Mime": "audio\/vnd.rip",
"Ext": ".rip"
},
{
"Name": "Open Web Media Project - Audio",
"Mime": "audio\/webm",
"Ext": ".weba"
},
{
"Name": "Advanced Audio Coding (AAC)",
"Mime": "audio\/x-aac",
"Ext": ".aac"
},
{
"Name": "Audio Interchange File Format",
"Mime": "audio\/x-aiff",
"Ext": ".aif"
},
{
"Name": "M3U (Multimedia Playlist)",
"Mime": "audio\/x-mpegurl",
"Ext": ".m3u"
},
{
"Name": "Microsoft Windows Media Audio Redirector",
"Mime": "audio\/x-ms-wax",
"Ext": ".wax"
},
{
"Name": "Microsoft Windows Media Audio",
"Mime": "audio\/x-ms-wma",
"Ext": ".wma"
},
{
"Name": "Real Audio Sound",
"Mime": "audio\/x-pn-realaudio",
"Ext": ".ram"
},
{
"Name": "Real Audio Sound",
"Mime": "audio\/x-pn-realaudio-plugin",
"Ext": ".rmp"
},
{
"Name": "Waveform Audio File Format (WAV)",
"Mime": "audio\/x-wav",
"Ext": ".wav"
},
{
"Name": "ChemDraw eXchange file",
"Mime": "chemical\/x-cdx",
"Ext": ".cdx"
},
{
"Name": "Crystallographic Interchange Format",
"Mime": "chemical\/x-cif",
"Ext": ".cif"
},
{
"Name": "CrystalMaker Data Format",
"Mime": "chemical\/x-cmdf",
"Ext": ".cmdf"
},
{
"Name": "Chemical Markup Language",
"Mime": "chemical\/x-cml",
"Ext": ".cml"
},
{
"Name": "Chemical Style Markup Language",
"Mime": "chemical\/x-csml",
"Ext": ".csml"
},
{
"Name": "XYZ File Format",
"Mime": "chemical\/x-xyz",
"Ext": ".xyz"
},
{
"Name": "Bitmap Image File",
"Mime": "image\/bmp",
"Ext": ".bmp"
},
{
"Name": "Computer Graphics Metafile",
"Mime": "image\/cgm",
"Ext": ".cgm"
},
{
"Name": "G3 Fax Image",
"Mime": "image\/g3fax",
"Ext": ".g3"
},
{
"Name": "Graphics Interchange Format",
"Mime": "image\/gif",
"Ext": ".gif"
},
{
"Name": "Image Exchange Format",
"Mime": "image\/ief",
"Ext": ".ief"
},
{
"Name": "JPEG Image",
"Mime": "image\/jpeg",
"Ext": ".jpeg, .jpg"
},
{
"Name": "OpenGL Textures (KTX)",
"Mime": "image\/ktx",
"Ext": ".ktx"
},
{
"Name": "Portable Network Graphics (PNG)",
"Mime": "image\/png",
"Ext": ".png"
},
{
"Name": "BTIF",
"Mime": "image\/prs.btif",
"Ext": ".btif"
},
{
"Name": "Scalable Vector Graphics (SVG)",
"Mime": "image\/svg+xml",
"Ext": ".svg"
},
{
"Name": "Tagged Image File Format",
"Mime": "image\/tiff",
"Ext": ".tiff"
},
{
"Name": "Photoshop Document",
"Mime": "image\/vnd.adobe.photoshop",
"Ext": ".psd"
},
{
"Name": "DECE Graphic",
"Mime": "image\/vnd.dece.graphic",
"Ext": ".uvi"
},
{
"Name": "Close Captioning - Subtitle",
"Mime": "image\/vnd.dvb.subtitle",
"Ext": ".sub"
},
{
"Name": "DjVu",
"Mime": "image\/vnd.djvu",
"Ext": ".djvu"
},
{
"Name": "DWG Drawing",
"Mime": "image\/vnd.dwg",
"Ext": ".dwg"
},
{
"Name": "AutoCAD DXF",
"Mime": "image\/vnd.dxf",
"Ext": ".dxf"
},
{
"Name": "FastBid Sheet",
"Mime": "image\/vnd.fastbidsheet",
"Ext": ".fbs"
},
{
"Name": "FlashPix",
"Mime": "image\/vnd.fpx",
"Ext": ".fpx"
},
{
"Name": "FAST Search & Transfer ASA",
"Mime": "image\/vnd.fst",
"Ext": ".fst"
},
{
"Name": "EDMICS 2000",
"Mime": "image\/vnd.fujixerox.edmics-mmr",
"Ext": ".mmr"
},
{
"Name": "EDMICS 2000",
"Mime": "image\/vnd.fujixerox.edmics-rlc",
"Ext": ".rlc"
},
{
"Name": "Microsoft Document Imaging Format",
"Mime": "image\/vnd.ms-modi",
"Ext": ".mdi"
},
{
"Name": "FlashPix",
"Mime": "image\/vnd.net-fpx",
"Ext": ".npx"
},
{
"Name": "WAP Bitamp (WBMP)",
"Mime": "image\/vnd.wap.wbmp",
"Ext": ".wbmp"
},
{
"Name": "eXtended Image File Format (XIFF)",
"Mime": "image\/vnd.xiff",
"Ext": ".xif"
},
{
"Name": "WebP Image",
"Mime": "image\/webp",
"Ext": ".webp"
},
{
"Name": "CMU Image",
"Mime": "image\/x-cmu-raster",
"Ext": ".ras"
},
{
"Name": "Corel Metafile Exchange (CMX)",
"Mime": "image\/x-cmx",
"Ext": ".cmx"
},
{
"Name": "FreeHand MX",
"Mime": "image\/x-freehand",
"Ext": ".fh"
},
{
"Name": "Icon Image",
"Mime": "image\/x-icon",
"Ext": ".ico"
},
{
"Name": "PCX Image",
"Mime": "image\/x-pcx",
"Ext": ".pcx"
},
{
"Name": "PICT Image",
"Mime": "image\/x-pict",
"Ext": ".pic"
},
{
"Name": "Portable Anymap Image",
"Mime": "image\/x-portable-anymap",
"Ext": ".pnm"
},
{
"Name": "Portable Bitmap Format",
"Mime": "image\/x-portable-bitmap",
"Ext": ".pbm"
},
{
"Name": "Portable Graymap Format",
"Mime": "image\/x-portable-graymap",
"Ext": ".pgm"
},
{
"Name": "Portable Pixmap Format",
"Mime": "image\/x-portable-pixmap",
"Ext": ".ppm"
},
{
"Name": "Silicon Graphics RGB Bitmap",
"Mime": "image\/x-rgb",
"Ext": ".rgb"
},
{
"Name": "X BitMap",
"Mime": "image\/x-xbitmap",
"Ext": ".xbm"
},
{
"Name": "X PixMap",
"Mime": "image\/x-xpixmap",
"Ext": ".xpm"
},
{
"Name": "X Window Dump",
"Mime": "image\/x-xwindowdump",
"Ext": ".xwd"
},
{
"Name": "Email Message",
"Mime": "message\/rfc822",
"Ext": ".eml"
},
{
"Name": "Initial Graphics Exchange Specification (IGES)",
"Mime": "model\/iges",
"Ext": ".igs"
},
{
"Name": "Mesh Data Type",
"Mime": "model\/mesh",
"Ext": ".msh"
},
{
"Name": "COLLADA",
"Mime": "model\/vnd.collada+xml",
"Ext": ".dae"
},
{
"Name": "Autodesk Design Web Format (DWF)",
"Mime": "model\/vnd.dwf",
"Ext": ".dwf"
},
{
"Name": "Geometric Description Language (GDL)",
"Mime": "model\/vnd.gdl",
"Ext": ".gdl"
},
{
"Name": "Gen-Trix Studio",
"Mime": "model\/vnd.gtw",
"Ext": ".gtw"
},
{
"Name": "Virtue MTS",
"Mime": "model\/vnd.mts",
"Ext": ".mts"
},
{
"Name": "Virtue VTU",
"Mime": "model\/vnd.vtu",
"Ext": ".vtu"
},
{
"Name": "Virtual Reality Modeling Language",
"Mime": "model\/vrml",
"Ext": ".wrl"
},
{
"Name": "iCalendar",
"Mime": "text\/calendar",
"Ext": ".ics"
},
{
"Name": "Cascading Style Sheets (CSS)",
"Mime": "text\/css",
"Ext": ".css"
},
{
"Name": "Comma-Seperated Values",
"Mime": "text\/csv",
"Ext": ".csv"
},
{
"Name": "HyperText Markup Language (HTML)",
"Mime": "text\/html",
"Ext": ".html"
},
{
"Name": "Notation3",
"Mime": "text\/n3",
"Ext": ".n3"
},
{
"Name": "Text File",
"Mime": "text\/plain",
"Ext": ".txt"
},
{
"Name": "PRS Lines Tag",
"Mime": "text\/prs.lines.tag",
"Ext": ".dsc"
},
{
"Name": "Rich Text Format (RTF)",
"Mime": "text\/richtext",
"Ext": ".rtx"
},
{
"Name": "Standard Generalized Markup Language (SGML)",
"Mime": "text\/sgml",
"Ext": ".sgml"
},
{
"Name": "Tab Seperated Values",
"Mime": "text\/tab-separated-values",
"Ext": ".tsv"
},
{
"Name": "troff",
"Mime": "text\/troff",
"Ext": ".t"
},
{
"Name": "Turtle (Terse RDF Triple Language)",
"Mime": "text\/turtle",
"Ext": ".ttl"
},
{
"Name": "URI Resolution Services",
"Mime": "text\/uri-list",
"Ext": ".uri"
},
{
"Name": "Curl - Applet",
"Mime": "text\/vnd.curl",
"Ext": ".curl"
},
{
"Name": "Curl - Detached Applet",
"Mime": "text\/vnd.curl.dcurl",
"Ext": ".dcurl"
},
{
"Name": "Curl - Source Code",
"Mime": "text\/vnd.curl.scurl",
"Ext": ".scurl"
},
{
"Name": "Curl - Manifest File",
"Mime": "text\/vnd.curl.mcurl",
"Ext": ".mcurl"
},
{
"Name": "mod_fly \/ fly.cgi",
"Mime": "text\/vnd.fly",
"Ext": ".fly"
},
{
"Name": "FLEXSTOR",
"Mime": "text\/vnd.fmi.flexstor",
"Ext": ".flx"
},
{
"Name": "Graphviz",
"Mime": "text\/vnd.graphviz",
"Ext": ".gv"
},
{
"Name": "In3D - 3DML",
"Mime": "text\/vnd.in3d.3dml",
"Ext": ".3dml"
},
{
"Name": "In3D - 3DML",
"Mime": "text\/vnd.in3d.spot",
"Ext": ".spot"
},
{
"Name": "J2ME App Descriptor",
"Mime": "text\/vnd.sun.j2me.app-descriptor",
"Ext": ".jad"
},
{
"Name": "Wireless Markup Language (WML)",
"Mime": "text\/vnd.wap.wml",
"Ext": ".wml"
},
{
"Name": "Wireless Markup Language Script (WMLScript)",
"Mime": "text\/vnd.wap.wmlscript",
"Ext": ".wmls"
},
{
"Name": "Assembler Source File",
"Mime": "text\/x-asm",
"Ext": ".s"
},
{
"Name": "C Source File",
"Mime": "text\/x-c",
"Ext": ".c"
},
{
"Name": "Fortran Source File",
"Mime": "text\/x-fortran",
"Ext": ".f"
},
{
"Name": "Pascal Source File",
"Mime": "text\/x-pascal",
"Ext": ".p"
},
{
"Name": "Java Source File",
"Mime": "text\/x-java-source,java",
"Ext": ".java"
},
{
"Name": "Setext",
"Mime": "text\/x-setext",
"Ext": ".etx"
},
{
"Name": "UUEncode",
"Mime": "text\/x-uuencode",
"Ext": ".uu"
},
{
"Name": "vCalendar",
"Mime": "text\/x-vcalendar",
"Ext": ".vcs"
},
{
"Name": "vCard",
"Mime": "text\/x-vcard",
"Ext": ".vcf"
},
{
"Name": "3GP",
"Mime": "video\/3gpp",
"Ext": ".3gp"
},
{
"Name": "3GP2",
"Mime": "video\/3gpp2",
"Ext": ".3g2"
},
{
"Name": "H.261",
"Mime": "video\/h261",
"Ext": ".h261"
},
{
"Name": "H.263",
"Mime": "video\/h263",
"Ext": ".h263"
},
{
"Name": "H.264",
"Mime": "video\/h264",
"Ext": ".h264"
},
{
"Name": "JPGVideo",
"Mime": "video\/jpeg",
"Ext": ".jpgv"
},
{
"Name": "JPEG 2000 Compound Image File Format",
"Mime": "video\/jpm",
"Ext": ".jpm"
},
{
"Name": "Motion JPEG 2000",
"Mime": "video\/mj2",
"Ext": ".mj2"
},
{
"Name": "MPEG-4 Video",
"Mime": "video\/mp4",
"Ext": ".mp4"
},
{
"Name": "MPEG Video",
"Mime": "video\/mpeg",
"Ext": ".mpeg"
},
{
"Name": "Ogg Video",
"Mime": "video\/ogg",
"Ext": ".ogv"
},
{
"Name": "Quicktime Video",
"Mime": "video\/quicktime",
"Ext": ".qt"
},
{
"Name": "DECE High Definition Video",
"Mime": "video\/vnd.dece.hd",
"Ext": ".uvh"
},
{
"Name": "DECE Mobile Video",
"Mime": "video\/vnd.dece.mobile",
"Ext": ".uvm"
},
{
"Name": "DECE PD Video",
"Mime": "video\/vnd.dece.pd",
"Ext": ".uvp"
},
{
"Name": "DECE SD Video",
"Mime": "video\/vnd.dece.sd",
"Ext": ".uvs"
},
{
"Name": "DECE Video",
"Mime": "video\/vnd.dece.video",
"Ext": ".uvv"
},
{
"Name": "FAST Search & Transfer ASA",
"Mime": "video\/vnd.fvt",
"Ext": ".fvt"
},
{
"Name": "MPEG Url",
"Mime": "video\/vnd.mpegurl",
"Ext": ".mxu"
},
{
"Name": "Microsoft PlayReady Ecosystem Video",
"Mime": "video\/vnd.ms-playready.media.pyv",
"Ext": ".pyv"
},
{
"Name": "DECE MP4",
"Mime": "video\/vnd.uvvu.mp4",
"Ext": ".uvu"
},
{
"Name": "Vivo",
"Mime": "video\/vnd.vivo",
"Ext": ".viv"
},
{
"Name": "Open Web Media Project - Video",
"Mime": "video\/webm",
"Ext": ".webm"
},
{
"Name": "Flash Video",
"Mime": "video\/x-f4v",
"Ext": ".f4v"
},
{
"Name": "FLI\/FLC Animation Format",
"Mime": "video\/x-fli",
"Ext": ".fli"
},
{
"Name": "Flash Video",
"Mime": "video\/x-flv",
"Ext": ".flv"
},
{
"Name": "M4v",
"Mime": "video\/x-m4v",
"Ext": ".m4v"
},
{
"Name": "Microsoft Advanced Systems Format (ASF)",
"Mime": "video\/x-ms-asf",
"Ext": ".asf"
},
{
"Name": "Microsoft Windows Media",
"Mime": "video\/x-ms-wm",
"Ext": ".wm"
},
{
"Name": "Microsoft Windows Media Video",
"Mime": "video\/x-ms-wmv",
"Ext": ".wmv"
},
{
"Name": "Microsoft Windows Media Audio\/Video Playlist",
"Mime": "video\/x-ms-wmx",
"Ext": ".wmx"
},
{
"Name": "Microsoft Windows Media Video Playlist",
"Mime": "video\/x-ms-wvx",
"Ext": ".wvx"
},
{
"Name": "Audio Video Interleave (AVI)",
"Mime": "video\/x-msvideo",
"Ext": ".avi"
},
{
"Name": "SGI Movie",
"Mime": "video\/x-sgi-movie",
"Ext": ".movie"
},
{
"Name": "CoolTalk",
"Mime": "x-conference\/x-cooltalk",
"Ext": ".ice"
},
{
"Name": "BAS Partitur Format",
"Mime": "text\/plain-bas",
"Ext": ".par"
},
{
"Name": "YAML Aint Markup Language \/ Yet Another Markup Language",
"Mime": "text\/yaml",
"Ext": ".yaml"
}]
`)

type MIMETypes struct {
	Name string `json:"name,ommitempty"`
	Mime string `json:"mime,ommitempty"`
	Ext  string `json:"ext,ommitempty"`
}
