#define _USE_MATH_DEFINES
#include <iostream>
#include <curl/curl.h>
#include <nlohmann/json.hpp>
#include <ctime>
#include <chrono>
#include <thread>
#include <cmath>
#include <vector>


struct GroundUnit {
    std::string id;
    std::string type;
    std::string squadron;
    double lat;
    double lon;
    std::string status;
    std::string time_stamp;
};

std::vector<GroundUnit> units = {
    {"TACP-1", "OPERATOR", "7th_ASOS", 31.8457, -106.4309, "ACTIVE", ""},
    {"TACP-2", "OPERATOR", "7th_ASOS", 31.8521, -106.4401, "ACTIVE", ""},
    {"FURY-1", "TANK", "1st_Armored", 31.7823, -106.3201, "ACTIVE", ""},
    {"FURY-2", "TANK", "1st_Armored", 31.7901, -106.3389, "ACTIVE", ""},
    {"FURY-3", "TANK", "1st_Armored", 31.7765, -106.3102, "ACTIVE", ""},
    {"GHOST-1", "RECON", "75th_Rangers", 31.9102, -106.5001, "ACTIVE", ""},
    {"GHOST-2", "RECON", "75th_Rangers", 31.9234, -106.5123, "ACTIVE", ""},
    {"GHOST-3", "RECON", "75th_Rangers", 31.8998, -106.4889, "ACTIVE", ""},
    {"BRAVO-1", "INFANTRY", "82nd_Airborne", 31.8201, -106.3701, "ACTIVE", ""},
    {"BRAVO-2", "INFANTRY", "82nd_Airborne", 31.8312, -106.3812, "ACTIVE", ""},
    {"BRAVO-3", "INFANTRY", "82nd_Airborne", 31.8089, -106.3612, "ACTIVE", ""},
    {"VIPER-1", "AH-64", "101st_Airborne", 31.8701, -106.4501, "ACTIVE", ""},
    {"VIPER-2", "AH-64", "101st_Airborne", 31.8812, -106.4612, "ACTIVE", ""},
    {"EAGLE-1", "F-16", "20th_Fighter_Wing", 31.9401, -106.5301, "AIRBORNE", ""},
    {"EAGLE-2", "F-16", "20th_Fighter_Wing", 31.9512, -106.5412, "AIRBORNE", ""},
    {"WOLF-1", "RECON", "160th_SOAR", 31.7601, -106.2901, "ACTIVE", ""},
    {"WOLF-2", "RECON", "160th_SOAR", 31.7712, -106.3012, "ACTIVE", ""},
    {"STRIKER-1", "ARTILLERY", "75th_FA", 31.7401, -106.2501, "ACTIVE", ""},
    {"STRIKER-2", "ARTILLERY", "75th_FA", 31.7512, -106.2612, "ACTIVE", ""},
    {"STRIKER-3", "ARTILLERY", "75th_FA", 31.7289, -106.2389, "ACTIVE", ""},
};

struct Threat {
    std::string id;
    std::string description;
    double lat;
    double lon;
    std::string status;
    std::string time_stamp;
};

double movement(double curr_pos, double target_pos, double step) {
    if (curr_pos < target_pos) {
        return curr_pos + step;
    }
    return curr_pos - step;
}

size_t write_call_back(char* data, size_t size, size_t nmemb, std::string* response) {
    response->append(data, size * nmemb);
    return size * nmemb;
}

void put_request(GroundUnit unit) {
    // open a conn to curl
    CURL* curl = curl_easy_init();
    nlohmann::json j;
    j["Lat"] = unit.lat;
    j["Lon"] = unit.lon;
    j["Status"] = unit.status;
    j["TimeStamp"] = unit.time_stamp;

    std::string body = j.dump();
    std::string url = "http://localhost:8080/unit/" + unit.id;
    std::string response = "";
    struct curl_slist* headers = nullptr;
    headers = curl_slist_append(headers, "Content-Type: application/json");

    curl_easy_setopt(curl, CURLOPT_URL, url.c_str()); // SET URL
    curl_easy_setopt(curl, CURLOPT_CUSTOMREQUEST, "PUT"); // SET AS A PUT
    curl_easy_setopt(curl, CURLOPT_POSTFIELDS, body.c_str()); // SET BODY
    curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_call_back);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response);


    curl_easy_perform(curl);
    curl_easy_cleanup(curl);
    curl_slist_free_all(headers);
}

void post_threat(GroundUnit unit, Threat threat) {
    CURL* curl = curl_easy_init();
    nlohmann::json j;
    j["Unit"] = {{"ID", unit.id}, {"Lat", unit.lat},
        {"Lon", unit.lon}, {"Status", unit.status}, {"TimeStamp", unit.time_stamp}};
    j["Target"] = {{"ID", threat.id}, {"Description", threat.description}, {"Lat", threat.lat},
        {"Lon", threat.lon}, {"Status", threat.status}, {"TimeStamp", threat.time_stamp}};
    

    std::string body = j.dump();
    std::string url = "http://localhost:8080/report-threat";
    std::string response = "";
    struct curl_slist* headers = nullptr;
    headers = curl_slist_append(headers, "Content-Type: application/json");

    curl_easy_setopt(curl, CURLOPT_URL, url.c_str());
    curl_easy_setopt(curl, CURLOPT_POSTFIELDS, body.c_str());
    curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_call_back);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response);

    curl_easy_perform(curl);
    curl_easy_cleanup(curl);
    curl_slist_free_all(headers);
}



std::string get_request(std::string url) {
    CURL* curl = curl_easy_init();
    std::string response;

    curl_easy_setopt(curl, CURLOPT_URL, url.c_str());
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_call_back);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response);
    curl_easy_perform(curl);
    curl_easy_cleanup(curl);
    return response;
}

double haversine(double lat1, double lon1, double lat2, double lon2) {
    // Earth's radius in m,iles
    const double R = 3958.8;

    // Convert degrees to radians
    lat1 = lat1 * M_PI / 180;
    lon1 = lon1 * M_PI / 180;
    lat2 = lat2 * M_PI / 180;
    lon2 = lon2 * M_PI / 180;

    double dLat = lat2 - lat1;
    double dLon = lon2 - lon1;

    double a = sin(dLat/2)*sin(dLat/2) + cos(lat1)*cos(lat2)*sin(dLon/2)*sin(dLon/2);
    double c = 2 * atan2(sqrt(a), sqrt(1-a));

    return R * c;

}

void random_threat(GroundUnit unit, int counter) {
        std::string id = unit.id + "-THREAT-" + std::to_string(counter);
        Threat threat = {id, "TEST", unit.lat, unit.lon, "ACTIVE", unit.time_stamp};
        post_threat(unit, threat);
}


void run_unit(GroundUnit unit) {
    std::time_t now = std::time(nullptr);
    std::string ts = std::ctime(&now);
    ts.erase(ts.find('\n'));
    srand(std::time(nullptr) + std::hash<std::string>{}(unit.id));

    double start_lat = unit.lat;
    double start_lon = unit.lon;
    double angle = 0.0;
    int counter = 1;

    while (true) {
        unit.lat = start_lat + 0.01 * sin(angle);
        unit.lon = start_lon + 0.01 * cos(angle);
        angle += 0.1;
        unit.time_stamp = ts;
        
        if (rand() % 5 == 0) {
            unit.status = "ENGAGED";
            random_threat(unit, counter);
            ++counter;
            
        }

        put_request(unit);
        std::this_thread::sleep_for(std::chrono::seconds(5));
        unit.status = "ACTIVE";
    }
}


int main() {
    for (GroundUnit unit : units) {
        std::thread t(run_unit, unit);
        t.detach();
    }

    while (true) {
        std::this_thread::sleep_for(std::chrono::hours(24));
        
    }


    return 0;
}