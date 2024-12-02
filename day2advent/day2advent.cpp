#include <iostream>
#include<vector>
#include<fstream>
#include "LevelSet.h"
#include<string>

int main()
{    
    std::vector<LevelSet> levelSets;
    std::ifstream file("data.txt");

    std::string lineText;

    while (std::getline(file, lineText)) {

        std::string currentNum = "";

        LevelSet set;
        for (int i = 0; i < lineText.length(); i++) {
            if (lineText[i] == ' ' || lineText[i] == '\n') {
                
                if (currentNum != "") {
                    set.addLevel(std::stoi(currentNum));
                    currentNum = "";
                }

                continue;
            }

            currentNum += lineText[i];
        }

        if (currentNum != "") {
            set.addLevel(std::stoi(currentNum));
        }

        levelSets.push_back(set);

    }

    int safeSets = 0;
    for (int i = 0; i < levelSets.size(); i++) {
        levelSets[i].show(std::cout);

        if (levelSets[i].isSafeWithDampener()) {
            safeSets++;
        }
    }

    std::cout << "There are: " << safeSets << " safe sets" << std::endl;

}
