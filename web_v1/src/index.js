import DPlayer from 'dplayer'
var dp = new DPlayer({
    container:document.getElementById('dplayer'),
    video:{
        url:"",
    },
    screenshot:true,

});
(function localFileVideoPlayer() {
    var playSelectedFile = function (event) {
        var file = this.files[0]
        var fileURL = URL.createObjectURL(file)
        dp.video.src = fileURL
    }
    var inputNode = document.querySelector('input')
    inputNode.addEventListener('change', playSelectedFile, false)
})();