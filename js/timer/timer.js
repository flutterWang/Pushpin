

function bubbleSort(FF) {
    for(let i = FF.length; i > 0; i--) {
     for(let j = 0; j < i; j++) {
       if(FF[j] > FF[j + 1]) {
         var k;
         k = FF[j];
         FF[j] = FF[j + 1];
         FF[j + 1] = k;
       }
     }
   }
   return FF;
}

function selectSort(Arr) {
    let i = 0;
    for (let j = 0; j < Arr.length; j++) {
        i = j
        for(let n = j + 1; n < Arr.length; n++) {
            if (Arr[i] < Arr[n]) {
                continue
            } else {
                i = n
            }
        }
        var k;
        k = Arr[i]
        Arr[i] = Arr[j]
        Arr[j] = k
    }

    return Arr 
}

function insertSort(Arr) {
    for (let i = 1; i < Arr.length; i++) {
        let h = Arr[i]
        for (var j = i - 1; j >= 0 && Arr[j] > h; j--) {
            Arr[j + 1] = Arr[j]
        }
        Arr[j + 1] = h
    }

    return Arr
}

// function insertSort(Arr) {
//     let result = []
//     for (let i = 0; i < Arr.length; i++) {
//         if (result.length == 0) {
//             result.push(Arr[i])
//         } else {
//             if (Arr[i] >= result[result.length - 1]) {
//                 result.push(Arr[i])
//                 continue
//             }
//             for (let j = result.length; j > 0; j--) {
//                 if (Arr[i] >= result[j]) {
//                     continue
//                 } else {
//                     result.splice(j, 0, Arr[i])
//                 }
//             }    
//         } 
//         console.log(result)
//     }
    
//     return result
// }

function randomArr(number) {
    let arr = []
    for (let i = 0; i < number; i++) {
        arr.push(Math.random() * number)
    }

    return arr
}

function timerForSort(Arr, func) {
    let timestampStart = (new Date()).valueOf()

    func(Arr)

    let timestampEnd = (new Date()).valueOf()
    return timestampEnd - timestampStart
}

let arr = randomArr(100000)
console.log(`Generate random array succeed ! Arr length: ${arr.length}`)
console.log(timerForSort(arr, insertSort))
