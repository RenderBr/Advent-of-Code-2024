use std::fs;
use std::thread::current;

// find valid mul(x,y) sets

fn main() {
    let mut multis: Vec<MultiPair> = vec![];
    let mut currently_processing_string = String::from("");

    let supported_chars = ['m','u','l','(',')',',','d','o','n','t','\''];

    let input = fs::read_to_string("./data.txt").unwrap();
    let mut processing_multi = false;
    let mut dontMode = false;

    for c in input.chars(){
        // go through each char
        if(currently_processing_string.contains("don't")){
            dontMode = true;
        }else if(currently_processing_string.contains("do")){
            dontMode = false;
        }


        // is it a valid multi currently?
        if(inside_multi_opening(&currently_processing_string) &&
            currently_processing_string.ends_with(")")){

            let processed_multi = process_multi(&currently_processing_string);

            println!("Found multi pair: {:?}", &processed_multi);

            if(processed_multi.is_some() && !dontMode){
                multis.push(processed_multi.unwrap())
            }

            currently_processing_string = String::new();
        }

        if(c == 'm'){
            currently_processing_string.clear();
            processing_multi = true;
        }

        // is the character not part of supported chars AND not a number? skip
        if(!supported_chars.contains(&c) && c.to_digit(10) == None){
            if(processing_multi){
                currently_processing_string.clear();
                processing_multi = false;
            }

            continue;
        }

        // push valid char to processor
        currently_processing_string.push(c);

        // print for debugging
        println!("{}", currently_processing_string);
    }

    println!("{:?}", multis);

    let mut sum: i32 = 0;
    if(multis.len() > 0){
        for pair in multis{
            sum += (pair.left*pair.right);
        }
    }

    println!("Sum: {}", sum);
}

fn inside_multi_opening(str: &String) -> bool {
    if(str.starts_with("mul(")){
        return true;
    }else{
        return false;
    }
}

fn process_multi(multi_str: &String) -> Option<MultiPair> {
    let mut left: Option<i32> = None;
    let mut right: Option<i32> = None;

    let mut currentNumStr: String = String::new();

    for c in multi_str.chars(){
        if(c == ',' || c == ')'){
            if let Ok(_parsedNum) = currentNumStr.parse::<i32>(){
                if(left.is_none()){
                    left = Some(_parsedNum);
                    currentNumStr.clear();
                    continue;
                }

                if(right.is_none()){
                    right = Some(_parsedNum);
                    currentNumStr.clear();
                    continue;
                }
            }
        }

        if(c.to_digit(10) != None){
            currentNumStr.push(c);
        }
    }

    if(left.is_some() && right.is_some()){
        return Some(MultiPair {
            left: left.unwrap(),
            right: right.unwrap()
        });
    }else{
        return None;
    }
}

#[derive(Debug)]
struct MultiPair{
    left: i32,
    right: i32
}