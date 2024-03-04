window.onload = function () {
    // 创建一个新的div元素
    const newDiv = document.createElement("div");

    // 添加一些文本内容到新的div中
    const newText = document.createTextNode("这是新添加的内容！");
    newDiv.appendChild(newText);

    // 设置新div的样式
    newDiv.style.color = "red";
    newDiv.style.fontSize = "20px";

    // 将新div添加到body中
    document.body.appendChild(newDiv);

    const div = document.getElementById("info");
    div.textContent = "begin to post message...";
    try {
        window.webkit.messageHandlers.greet
            .postMessage("from web backend")
            .then(
                (info) => {
                    div.textContent = info;
                },
                (err) => {
                    div.textContent = "return error:" + err.message;
                },
            );
    } catch (err) {
        div.textContent = "error:" + err.message;
    }
};
