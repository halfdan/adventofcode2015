def groupAnagrams(items):
	agrams = {}

	for item in items:
		st = "".join(sorted(item))
		if st in agrams:
			agrams[st] = agrams[st] + [item]
		else:
			agrams[st] = [item]

	return agrams.values()

print(groupAnagrams(["eat","tea","ten","poop","net","ate"]))
