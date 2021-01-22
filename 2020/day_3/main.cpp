#include <iostream>
#include <fstream>
#include <string>
#include <vector>

int howManyTrees(int right, int down, std::vector<std::string> data){
    int trees = 0;
    int r = 0, c = 0;
    while(r < data.size()){
        std::string curr = data[r];

        if(curr[c] == '#'){
            trees++;
        }

        r += down;
        c += right;
        c = c % curr.length();
    }

    return trees;
}


int main(){

    std::string filename = "input.txt";
    
    std::string line;
    std::ifstream readInput(filename);
    std::vector<std::string> data;

    while(std::getline (readInput, line)){
        data.push_back(line);
    }

    std::cout << "Trees hit going 3 right and 1 down = " << howManyTrees(3,1,data) << std::endl;

    int values[] = {howManyTrees(1,1,data), howManyTrees(3,1,data), howManyTrees(5,1,data), howManyTrees(7,1,data), howManyTrees(1,2,data)};
    int totalMult = values[0] * values[1] * values[2] * values[3] * values[4];

    std::cout << "Total trees multiplied = " << totalMult << std::endl;

    return 0;
}