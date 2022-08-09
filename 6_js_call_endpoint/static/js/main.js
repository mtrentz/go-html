let number = document.getElementById("mynumber");

let increment = () => {
  number.innerHTML = parseInt(number.innerHTML) + 1;
};

let input = document.getElementById("myinput");
let funcKeyUp = () => {
  val = input.value;

  request = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      value: val,
    }),
  };

  // Send a http req to /reverse to reverse the string in Go
  r = fetch("/reverse", request)
    .then((response) => {
      return response.text();
    })
    .then((text) => {
      document.getElementById("reversedString").innerHTML = text;
    });
};
