#include <iostream>
#include <fstream>
#include <string>
#include <vector>

struct passwordData{
    int lowBound;
    int upBound;
    char letter;
    std::string password;
};

int numValidPasswords(std::vector<passwordData> data){
    int totalValid = 0;
    for(int x = 0; x < data.size(); x++){
        int charCount = 0;
        passwordData current = data[x];
        for(int y = 0; y < current.password.length(); y++){
            if(current.password[y] == current.letter){
                charCount++;
            }
        }
        if(charCount >= current.lowBound && charCount <= current.upBound){
            totalValid++;
        }
    }
    return totalValid;
}

int numNewPasswords(std::vector<passwordData> data){
    int totalValid = 0;

    for(int x = 0; x < data.size(); x++){
        passwordData current = data[x];

        int idx1 = current.lowBound - 1;
        int idx2 = current.upBound - 1;

        bool pos1let = false, pos2let = false;

        if(current.password[idx1] == current.letter){
            pos1let = true;
        }
        if(current.password[idx2] == current.letter){
            pos2let = true;
        }

        if((pos1let && !pos2let) || (!pos1let && pos2let)){
            totalValid++;
        }
    }

    return totalValid;
}

int main(){
    std::string filename = "input.txt";
    
    std::string line;
    std::ifstream readInput(filename);
    std::vector<passwordData> data;

    while(std::getline (readInput, line)){
        std::string currentCharSeq = "";
        passwordData newData;
        newData.letter = '*';

        bool wait = false;
        for(int x = 0; x < line.length(); x++){
            char ch = line[x];
            if(wait){
                wait = false;
                continue;
            }else if(ch == '-'){
                newData.lowBound = std::stoi(currentCharSeq);
                currentCharSeq = "";

            }else if(ch == ' '){
                if(newData.letter == '*'){
                    newData.upBound = std::stoi(currentCharSeq);
                    currentCharSeq = "";
                }
            }else if(ch == ':'){
                newData.letter = currentCharSeq[0]; // It should only be one value
                currentCharSeq = "";

            }else{
                currentCharSeq += ch;
            }
        }
        newData.password = currentCharSeq;

        data.push_back(newData);
    }

    std::cout << "Number valid passwords = " << numValidPasswords(data) << std::endl;
    std::cout << "Number new valid passwords = " << numNewPasswords(data) << std::endl;

    return 0;
}