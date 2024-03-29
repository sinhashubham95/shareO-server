# SHAREO Project - Golang Server

### Technologies Used
1. Go 1.12.6
2. Redis
3. GraphQL
4. Gin-Gonic

This project illustrates launching a GraphQL server along with the playground. It invokes the Google's custom search API. As the daily limit is low, the use of Redis comes into picture. The data from the custom search API is cached into Redis with a TTL of 10 days.

This project also illustrates an interesting way to handle configs in Go. It takes the config file path as a VM argument and adds a watcher to it. Whenever the file is updated, it basically reloads the configurations. It is just a minor implementation of the famous project [Viper](https://github.com/spf13/viper).

### Steps to Run (MacOS)
1. Install Go using the following command - ```brew install go```. If you already have Go installed, you can check the version using ```go version```. If the version **<= 1.12.6**, then upgrade using the following command ```brew upgrade go```.
2. Install Redis using the following command - ```brew install redis```.
3. Clone this project - ```git clone https://github.com/sinhashubham95/shareO-server.git```.
4. Move into the directory - ```cd shareO-server```.
5. Create a file named **.config.json** in the project directory - ```touch .config.json```.
6. Add the following configs in that **PORT**, **customSearchURL**, **customSearchKey**, **customSearchClient**.
7. Now run the application using the following command - ```go run main.go CONFIG_FILE=${config_folder_path}/.config.json```.
8. Browse the GraphQL playground - ```http://localhost:${PORT}/graphql```. Replace the port with the value provided in the config file.

### How to get the Google Custom Search Credentials?
Follow the steps at [Google Custom Search Documentation](https://developers.google.com/custom-search/docs/tutorial/introduction).
