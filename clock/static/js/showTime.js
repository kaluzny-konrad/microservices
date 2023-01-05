function showTime(){
    let time = getTime();

    document.getElementById("clockDisplay").innerText = time;
    document.getElementById("clockDisplay").textContent = time;
    
    setTimeout(showTime, 1000);
}

showTime();