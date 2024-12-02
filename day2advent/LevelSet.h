#pragma once
#include<vector>
#include<iostream>

class LevelSet {
private:
	std::vector<int> set;
public:
	std::vector<int> getSet() const { return set; }
	int getCount() const { return set.size(); }
	void addLevel(int num);
	void show(std::ostream&);
	bool isSafe(std::vector<int>) const;
	bool isSafe() const;
	bool isSafeWithDampener();
	LevelSet();
};