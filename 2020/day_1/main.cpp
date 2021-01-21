// main.cpp
// Advent of Code 2020 - Day 1

#include <iostream>
#include <fstream>
#include <string>
#include <unordered_map>
#include <vector>

void twoSum(std::string filename){
    std::string line;
    std::ifstream readInput(filename);

    std::unordered_map<int,int> mapVal;

    while(std::getline (readInput, line)){
        int value = std::stoi(line);
        int remain = 2020 - value;

        // Check if the remainder is in the map
        if(mapVal.find(value) == mapVal.end()){ // If not, put it in and pair it with the value
            mapVal.insert(std::make_pair(remain,value));
        }else{ // If it's in the map, it's the value you need
            int out = value * remain;
            std::cout << "Two values are: " << value << ", " << remain << std::endl;
            std::cout << "Two sum out value = " << out << std::endl;
            return;
        }
    }
}

struct valuePair{
    int first;
    int second;
};

void threeSum(std::string filename){
    std::string line;
    std::ifstream readInput(filename);
    std::vector<int> valueVec;
    while(std::getline(readInput,line)){ // Push everything into a vector to make access easier
        valueVec.push_back(std::stoi(line));
    }

    std::unordered_map<int,valuePair> valMap;
    int goal = 2020;

    for(int x = 0; x < valueVec.size(); x++){
        int xVal = valueVec[x];
        if(valMap.find(xVal) != valMap.end()){
            
            valuePair outVals = valMap[xVal];
            std::cout << "Three values are: " << xVal << ", " << outVals.first << ", " << outVals.second << std::endl;
            int out = xVal * outVals.first * outVals.second;
            std::cout << "Three sum out val = " << out << std::endl;
            return;
        }
        for(int y = x+1; y < valueVec.size(); y++){
            int yVal = valueVec[y];
            if(valMap.find(yVal) != valMap.end()){
                valuePair outVals = valMap[yVal];
                std::cout << "Three values are: " << yVal << ", " << outVals.first << ", " << outVals.second << std::endl;
                int out = yVal * outVals.first * outVals.second;
                std::cout << "Three sum out val = " << out << std::endl;
                return;
            }else{
                int remain = goal - (xVal + yVal);
                if(remain > 0){
                    valuePair insert;
                    insert.first = xVal;
                    insert.second = yVal;
                    valMap.insert(std::make_pair(remain,insert));
                }
            } 
        }
    }
}

int main(){
    std::string filename = "input.txt";
    twoSum(filename);
    threeSum(filename);
    return 0;
}