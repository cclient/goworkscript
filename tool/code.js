
var topname="Data";
var defaulttype="string";
var bson=true; //对应 mongodb 
var json=true; //http response 
var scheme=false; //http request->scheme 
var jdata={
    "_id" : "564d5162e54b3106fb7badea",
    "macs" : [
        "00-21-26-00-C8-B0"
    ],
    "time" : 1447907400,
    "timestr" : "2015-11-19 12:30",
    "shop":{
        "name":"shop1"
    }
};
jdata={
          "sum": 1000,
          "data": [
              {
                  "province": "",
                  "city": "",
                  "gender": "",
                  "bankuaiid": "40",
                  "author": "Luke_Li",
                  "title": "2017骞�4鏈�7鏃ョ鍒拌褰曡创",
                  "url": "http://www.oneplusbbs.com/thread-3329370-407-1.html",
                  "comment_num": "",
                  "floor": 8129,
                  "view_num": "",
                  "content": "",
                  "time": 20170407,
                  "threadid": "a78ae9b3efdcc8ed9f7761a40a5797d7",
                  "device": "",
                  "bankuainame": "绀惧尯浜嬪姟",
                  "level": "",
                  "authorurl": "http://www.oneplusbbs.com/home.php?mod=space&uid=1395573"
              },
              {
                  "province": "",
                  "city": "",
                  "gender": "",
                  "bankuaiid": "40",
                  "author": "PicassoX",
                  "title": "2017骞�4鏈�7鏃ョ鍒拌褰曡创",
                  "url": "http://www.oneplusbbs.com/thread-3329370-439-1.html",
                  "comment_num": "",
                  "floor": 8764,
                  "view_num": "",
                  "content": "",
                  "time": 20170407,
                  "threadid": "f90005a81a1ae3895fcf9580dc9e55ff",
                  "device": "",
                  "bankuainame": "绀惧尯浜嬪姟",
                  "level": "",
                  "authorurl": "http://www.oneplusbbs.com/home.php?mod=space&uid=1167206"
              }
          ]
      }
String.prototype.firstToUpperCase=function(){
    return this[0].toUpperCase()+this.substring(1);
}
var fun=(function(){
    var otherobj=[];
    var goobjstring="";
    function getStruct(data,collectionname){
        goobjstring+="type "+collectionname.firstToUpperCase()+" struct {\n";
        var per="\t";
        for(var key in data){
            var newkey=key.firstToUpperCase();
            goobjstring+=per +newkey+" "+getType(data[key],key);
            if (json||bson||scheme){
                goobjstring+=' `';
                var temparr=[]
                if (json){
                    temparr.push('json:"'+key+'"');
                }
                if (bson){
                    temparr.push('bson:"'+key+'"');
                }
                if (scheme){
                    temparr.push('scheme:"'+key+'"');
                }
                goobjstring+=temparr.join(" ");
                goobjstring+='`';
            }
            goobjstring+="\n";
        }
        goobjstring+="}\n";
        while (otherobj.length>0){
            var subobj=otherobj.pop();
            getStruct(subobj.obj,subobj.key)
        }
        return goobjstring
    }
    function getType(obj,key){
        var type=defaulttype;
        if(obj){
            switch(obj.constructor)
            {
                case Array:
                    type="[]"+getType(obj[0]||"",key.firstToUpperCase()) ;
                    break;
                case Object:
                    otherobj.push({key:key,obj:obj});
                    type=key.firstToUpperCase()
                    break;
                case String:
                    type="string"
                    break;
                case Number:
                    type="int"
                    break;
                case Boolean:
                    type="bool"
                    break;
                default :
            }
        }
        return type;
    }
    return getStruct
})()

console.log(fun(jdata,topname))
