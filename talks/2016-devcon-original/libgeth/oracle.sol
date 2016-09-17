contract ReleaseOracle {
	function currentVersion() constant returns (uint32 major, uint32 minor, uint32 patch) { // HL
		// [...] Look up the currently active release
		return (release.major, release.minor, release.patch);
	}
	// [...]
}
