#define _USE_MATH_DEFINES
#include <iostream>
#include <curl/curl.h>
#include <ctime>
#include <thread>
#include <chrono>
#include <cstdlib>
#include <nlohmann/json.hpp>
#include <cmath>


size_t writeCallBack(char* data, size_t size, size_t nmemb, std::string* response) {
    response->append(data, size * nmemb);
    return size * nmemb;
}

std::string getRequest(std::string url) {
    CURL* curl = curl_easy_init();
    std::string response;

    curl_easy_setopt(curl, CURLOPT_URL, url.c_str());
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, writeCallBack);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response);
    curl_easy_perform(curl);
    curl_easy_cleanup(curl);
    return response;
}

void putRequest(std::string url) {
    CURL* curl = curl_easy_init();
    std::string response;

    curl_easy_setopt(curl, CURLOPT_URL, url.c_str());
    curl_easy_setopt(curl, CURLOPT_CUSTOMREQUEST, "PUT");
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, writeCallBack);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response);
    curl_easy_perform(curl);
    curl_easy_cleanup(curl);
    
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

struct Target {
    std::string ID;
    double Lat;
    double Lon;
};

Target getActiveTarget() {
    // Polls GET /targets
    Target activeTarget;
    
    bool active = false;
    while (!active) {
        std::cout << "POLLING FOR TARGET" << std::endl;
        std::string response = getRequest("http://localhost:8080/targets");
        if (!response.empty()) {
            auto targets = nlohmann::json::parse(response);
            for (auto target : targets) {
                if (target["Status"] == "ACTIVE") {
                    activeTarget.Lat = target["Lat"].get<double>();
                    activeTarget.Lon= target["Lon"].get<double>();
                    activeTarget.ID = target["ID"].get<std::string>();
                    active = true;
                    break;
                }
            }
        }
        std::this_thread::sleep_for(std::chrono::seconds(5));
    }
    return activeTarget;
}

void fireMission(Target target) {
    std::cout<< "CLEARED HOT - FIRING IN 3 SECONDS" << std::endl;
    for (int i = 3; i > 0; i--) {
        std::this_thread::sleep_for(std::chrono::seconds(1));
        std::cout << i << std::endl;
    }
    std::cout << "FIRE!!" << std::endl;
    std::this_thread::sleep_for(std::chrono::seconds(3));

    std::string url = "http://localhost:8080/target/" + target.ID + "?status=DESTROYED";
    putRequest(url);
}

void deconfliction(Target target) {
    bool clearedHot = false;
    double lat = target.Lat;
    double lon = target.Lon;
    double radius = .5;

    while (!clearedHot) {
        std::string url = "http://localhost:8080/units/nearby?lat=" +
                        std::to_string(lat) +
                        "&lon=" + std::to_string(lon) +
                        "&radius=" + std::to_string(radius);
        std::string response = getRequest(url);
        auto units = nlohmann::json::parse(response);
        if (units.empty()) {
            clearedHot = true;
            fireMission(target);
        } else {
            std::cout<<"HOLDING FIRE NEARBY UNITS"<<std::endl;
        }
        std::this_thread::sleep_for(std::chrono::seconds(5));
    }
}




int main() {
    Target target = getActiveTarget();
    deconfliction(target);
    return 0;
}