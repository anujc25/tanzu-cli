CREATE TABLE IF NOT EXISTS "PluginBinaries" (
		"PluginName"         TEXT NOT NULL,
		"Target"             TEXT NOT NULL,
		"RecommendedVersion" TEXT NOT NULL,
		"Version"            TEXT NOT NULL,
		"Hidden"             INTEGER NOT NULL,
		"Description"        TEXT NOT NULL,
		"Publisher"          TEXT NOT NULL,
		"Vendor"             TEXT NOT NULL,
		"OS"                 TEXT NOT NULL,
		"Architecture"       TEXT NOT NULL,
		"Digest"             TEXT NOT NULL,
		"URI"                TEXT NOT NULL,
		PRIMARY KEY("PluginName", "Target", "Version", "OS", "Architecture")
);