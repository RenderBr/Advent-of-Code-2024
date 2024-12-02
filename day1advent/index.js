const puzzleInput = "";

function parseAsLists(string){
    let list1 = [], list2 = [];
    
    let currentNumberString = "";
    let switchList = false;
    
    for (let charIndex = 0; charIndex < string.length; charIndex++){
        if(string[charIndex] === ' ' || string[charIndex] === '\n'){
            if (currentNumberString !== ""){
                
                if (switchList){
                    list1.push(Number(currentNumberString));    
                }else{
                    list2.push(Number(currentNumberString));
                }
                
                switchList = !switchList;
                currentNumberString = "";
            }
            continue;
        }
        
        currentNumberString += string[charIndex];
    }
    
    if (currentNumberString !== ""){
        if (switchList){
            list1.push(Number(currentNumberString));
        }else{
            list2.push(Number(currentNumberString));
        }
    }
    
    return {x:list1, y:list2};
}

function selectionSort(array){
    // Iterate through indexes
    let smallestIndex;
    for (let i = 0; i < array.length; i++) {
        smallestIndex = smallestOnOffset(array, i, i);
        let currentSwapValue = array[i];

        array[i] = array[smallestIndex];
        array[smallestIndex] = currentSwapValue;
    }
    
    return array;
}

function smallestOnOffset(array, offset, currentSmallestIndex){
    let smallest = currentSmallestIndex;
    
    for (let i = offset; i <array.length; i++){
        if (array[i] < array[smallest]){
            smallest = i;
        }
        
    }
    
    return smallest;
}

function findDistanceSum(array1, array2){
    let sum = 0;
    
    for (let i = 0; i < array1.length; i++){
        if (array1[i] !== array2[i]){
            sum += array1[i] > array2[i] ? (array1[i] - array2[i])  : array2[i] - array1[i];
        }
    }
    
    return sum;
}

function findSimilarityScore(array1, array2){
    let simScore = 0;

    for (let i = 0; i < array1.length; i++){
        simScore += getScore(array1[i],array2);
    }
    
    return simScore;
}

function getScore(num, array2){
    let count = 0;
    
    for (let i = 0; i < array2.length; i++){
        if (array2[i] === num){
            count++;
        }
    }
    
    return count*num;
}

let lists = parseAsLists(puzzleInput);

selectionSort(lists.x);
selectionSort(lists.y);

console.log(lists.x,lists.y)

const sum = findDistanceSum(lists.x,lists.y);

console.log(`Distance: ${sum}`);

const simScore = findSimilarityScore(lists.x,lists.y);
console.log(`Similarity Score: ${simScore}`);