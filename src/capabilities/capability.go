package capabilities

// Capability represents a feature capability
type Capability int32

// Capabilities are defined in devmgr.v1125api15.symbol.Capability
const (
	// public static final int CAPABILITY_NONE = 0;
	None Capability = 0

	// public static final int CAPABILITY_SHARED_VOLUME = 1;

	// public static final int CAPABILITY_STORAGE_POOLS_TO_4 = 2;

	// public static final int CAPABILITY_MIXED_RAIDLEVEL = 3;

	// public static final int CAPABILITY_AUTO_CODE_SYNC = 4;

	// public static final int CAPABILITY_AUTO_LUN_TRANSFER = 5;

	// public static final int CAPABILITY_SUB_LUNS_ALLOWED = 6;

	// public static final int CAPABILITY_STORAGE_POOLS_TO_8 = 7;

	// public static final int CAPABILITY_STORAGE_POOLS_TO_2 = 8;

	// public static final int CAPABILITY_STORAGE_POOLS_TO_MAX = 9;

	// public static final int CAPABILITY_STORAGE_POOLS_TO_64 = 10;

	// public static final int CAPABILITY_STORAGE_POOLS_TO_16 = 11;

	// public static final int CAPABILITY_SNAPSHOTS = 12;
	SnapshotVirtualDisks Capability = 12

	// public static final int CAPABILITY_REMOTE_MIRRORING = 13;

	// public static final int CAPABILITY_VOLUME_COPY = 14;
	VirtualDiskCopy Capability = 14

	// public static final int CAPABILITY_STAGED_DOWNLOAD = 15;

	// public static final int CAPABILITY_MIXED_DRIVE_TYPES = 16;

	// public static final int CAPABILITY_GOLD_KEY = 17;

	// public static final int CAPABILITY_DRIVE_TRAY_EXPANSION = 18;

	// public static final int CAPABILITY_BUNDLE_MIGRATION = 19;

	// public static final int CAPABILITY_STORAGE_POOLS_TO_128 = 20;

	// public static final int CAPABILITY_STORAGE_POOLS_TO_256 = 21;

	// public static final int CAPABILITY_RAID6 = 22;

	// public static final int CAPABILITY_PERFORMANCE_TIER = 23;
	HighPerformanceTier Capability = 23

	// public static final int CAPABILITY_STORAGE_POOLS_TO_32 = 24;

	// public static final int CAPABILITY_STORAGE_POOLS_TO_96 = 25;

	// public static final int CAPABILITY_STORAGE_POOLS_TO_192 = 26;

	// public static final int CAPABILITY_STORAGE_POOLS_TO_512 = 27;

	// public static final int CAPABILITY_REMOTE_MIRRORS_TO_16 = 28;

	// public static final int CAPABILITY_REMOTE_MIRRORS_TO_32 = 29;

	// public static final int CAPABILITY_REMOTE_MIRRORS_TO_64 = 30;

	// public static final int CAPABILITY_REMOTE_MIRRORS_TO_128 = 31;

	// public static final int CAPABILITY_SNAPSHOTS_PER_VOL_TO_4 = 32;

	// public static final int CAPABILITY_SNAPSHOTS_PER_VOL_TO_8 = 33;

	// public static final int CAPABILITY_SNAPSHOTS_PER_VOL_TO_16 = 34;

	// public static final int CAPABILITY_SNAPSHOTS_PER_VOL_TO_2 = 35;

	// public static final int CAPABILITY_SECURE_VOLUME = 36;

	// public static final int CAPABILITY_PROTECTION_INFORMATION = 37;
	DataAssurance Capability = 37

	// public static final int CAPABILITY_SSD_SUPPORT = 38;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_112 = 45;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_120 = 46;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_256 = 47;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_448 = 48;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_480 = 49;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_MAX = 50;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT = 51;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_12 = 53;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_16 = 54;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_24 = 55;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_32 = 56;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_48 = 57;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_60 = 58;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_64 = 59;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_72 = 60;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_96 = 61;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_128 = 62;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_136 = 63;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_144 = 64;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_180 = 65;
	PhysicalDiskSlots180 Capability = 65

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_192 = 66;
	PhysicalDiskSlots192 Capability = 66

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_272 = 67;

	// public static final int CAPABILITY_FDE_PROXY_KEY_MANAGEMENT = 68;

	// public static final int CAPABILITY_REMOTE_MIRRORS_TO_8 = 84;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_384 = 93;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_300 = 94;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_360 = 95;

	// public static final int CAPABILITY_FLASH_READ_CACHE = 96;
	SsdCache Capability = 96

	// public static final int CAPABILITY_STORAGE_POOLS_TYPE2 = 105;

	// public static final int CAPABILITY_REMOTE_MIRRORING_TYPE2 = 107;

	// public static final int CAPABILITY_TOTAL_NUMBER_OF_ARVM_MIRRORS_PER_ARRAY = 109;

	// public static final int CAPABILITY_TOTAL_NUMBER_OF_PITS_PER_ARRAY = 110;
	Snapshot Capability = 110

	// public static final int CAPABILITY_TOTAL_NUMBER_OF_THIN_VOLUMES_PER_ARRAY = 111;

	// public static final int CAPABILITY_DRIVE_SLOT_LIMIT_TO_240 = 113;

	// public static final int CAPABILITY_SNAPSHOTS_TYPE2 = 114;

	// public static final int CAPABILITY_TARGET_PORT_LUN_MAPPING = 115;
)
