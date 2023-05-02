import axios from 'axios'

const api = axios.create({
    //環境ごとに異なる値が自動的に設定されるP162
    baseURL: "http://localhost:8080",
    timeout: 5000,
    //timeout: 1000000,
    headers: {
      'Content-Type': 'application/json',
      //"X-CSRFToken": "ここに追加"
      //'X-Requested-With': 'XMLHttpRequest'
    }
  })

  //機密情報許可
  //POST
  api.interceptors.request.use(
    function(config) {
      // Do something before request is sent
      config.withCredentials = true;
      return config;
    },
    function(error) {
      // Do something with request error
      return Promise.reject(error);
    }
  );

  //GET
  api.interceptors.response.use(
    function(config) {
      // Do something before request is sent
      config.withCredentials = true;
      return config;
    },
    function(error) {
      // Do something with request error
      return Promise.reject(error);
    }
  );

  

  export default api