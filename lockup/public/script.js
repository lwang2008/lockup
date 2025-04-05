//let input = document.getElementById("data")
let submit = document.getElementById("submit")
let finalData = document.getElementById("finalData")
const ethereumButton = document.getElementById('connect');

let method = "personal_sign"

let sigData = "";

ethereumButton.addEventListener('click', () => {
  ethereum.request({ method: 'eth_requestAccounts' });
});


fetch("./getSignatureData").then(response => response.json()).then(data => {sigData = data})

submit.onclick= async () => {
  
  let address = ethereum.selectedAddress
  let params = [sigData.value,address]


  await ethereum.sendAsync(
    {
      method,
      params,
      address
    }, (err,result)  => { 
      if (err) {
        console.log(result)
      }
      
      finalData.innerHTML = (result.result) //.slice(0,-2)

      const options = {
        method: "POST",
        body: result.result
      }

      fetch("./submitSignature", options)


    })
}