#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <set>

class Passport{
public:
    bool byr, iyr, eyr, hgt, hcl, ecl, pid, cid;
    Passport(){
        this->byr = false;
        this->iyr = false;
        this->eyr = false;
        this->hgt = false;
        this->hcl = false;
        this->ecl = false;
        this->pid = false;
        this->cid = false;
    };
    ~Passport(){};
    bool isValid(){
        return (byr && iyr && eyr && hgt && hcl && ecl && pid);
    }
};


int validPassports(std::vector<std::string> data){
    int numPassport = 0;

    // cid is optional
    std::string byr = "byr", iyr = "iyr", eyr = "eyr", hgt = "hgt", hcl = "hcl", ecl = "ecl", pid = "pid", cid = "cid";

    Passport *passport = new Passport();

    for(int x = 0; x < data.size(); x++){
        std::string current = data[x];
        if(current == ""){ 
            if(passport->isValid()){
                numPassport++;
            }
            passport = new Passport(); // Reset the passport
        }
        for(int idx = 0; idx < current.length(); idx++){
            if(current[idx] == ':'){
                std::string field = current.substr(idx-3,3);
                if(field == byr) passport->byr = true;
                else if(field == iyr) passport->iyr = true;
                else if(field == eyr) passport->eyr = true;
                else if(field == hgt) passport->hgt = true;
                else if(field == hcl) passport->hcl = true;
                else if(field == ecl) passport->ecl = true;
                else if(field == pid) passport->pid = true;
                else passport->cid = true;
            }
        }
    }
    delete passport;
    return numPassport;
}

int dataValidation(std::vector<std::string> data){
    int numPassports = 0;
    std::string byr = "byr", iyr = "iyr", eyr = "eyr", hgt = "hgt", hcl = "hcl", ecl = "ecl", pid = "pid", cid = "cid";

    Passport *passport = new Passport();

    for(int x = 0; x < data.size(); x++){
        std::string current = data[x];
        if(current == ""){ 
            if(passport->isValid()){
                numPassports++;
            }
            passport = new Passport(); // Reset the passport
        }
        for(int idx = 0; idx < current.length(); idx++){
            if(current[idx] == ':'){
                std::string field = current.substr(idx-3,3);
                if(field == byr) {
                    int year = std::stoi(current.substr(idx+1,4));
                    if(year >= 1920 && year <= 2002){
                        passport->byr = true;
                    }
                }
                else if(field == iyr){
                    int year = std::stoi(current.substr(idx+1,4));
                    if(year >= 2010 && year <= 2020){
                        passport->iyr = true;
                    }
                }
                else if(field == eyr) {
                    int year = std::stoi(current.substr(idx+1,4));
                    if(year >= 2020 && year <= 2030){
                        passport->eyr = true;
                    }
                }
                else if(field == hgt) {
                    if(current.substr(idx+3,2) == "in"){

                    }else if(current.substr(idx+4,2) == "cm"){
                        int height = std::stoi(current.substr(idx+1,3))
                    }

                    passport->hgt = true;
                }
                else if(field == hcl) {
                    passport->hcl = true;
                }
                else if(field == ecl) {
                    passport->ecl = true;
                }
                else if(field == pid) {
                    passport->pid = true;
                }
                else {
                    passport->cid = true;
                }
            }
        }
    }
    delete passport;

    return numPassports;
}


int main(){

    std::string filename = "input.txt";
    
    std::string line;
    std::ifstream readInput(filename);
    std::vector<std::string> data;

    while(std::getline (readInput, line)){
        data.push_back(line);
    }

    std::cout << "Valid passports = " << validPassports(data) << std::endl;

    return 0;
}