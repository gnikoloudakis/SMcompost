
onmessage = function (e) {
    // console.log(e.data);
    var data =[];
    for (var i = 0; i < e.data.length; i++){
        // console.log(e.data[i]["Temperature"]);
        data.push([new Date(e.data[i]["Timestamp"]).getTime(), e.data[i]["Temperature"]])
    }
    postMessage(data);
};