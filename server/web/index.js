// 1.
axios.get("/apis/hello").then((resp) => {
    document.getElementById("result1").innerText = resp.data
}).catch()

// 2.
const url = (location.protocol === "https:" ? "wss:" : "ws:") + `//${window.location.host}/apis/ws`
console.log(url)
const sc = new WebSocket(url)
sc.onmessage = (e) => {
    document.getElementById("result2").innerText = e.data
}