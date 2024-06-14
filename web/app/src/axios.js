import ax from 'axios'
export const baseAPIURL = "http://127.0.0.1:8080/api"
export const axios = ax.create({
    baseURL: baseAPIURL,
    headers:{
        'Cache-Control': 'no-cache',
        'Pragma': 'no-cache',
        'Expires': '0',
        'Access-Control-Allow-Credentials':true
    }
})