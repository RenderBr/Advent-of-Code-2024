import java.sql.Array;
import java.util.*;

public class Update {
    public ArrayList<Integer> updates;
    public HashMap<Integer, ArrayList<Integer>> rules;

    public int middleNumber = 0;

    public Update(HashMap<Integer, ArrayList<Integer>> rules){
        updates = new ArrayList<>();
        this.rules = rules;
    }

    public Integer[] fix(){
        Integer[] fixedArray = updates.toArray(new Integer[0]);
        while(!isMatchingRules(new ArrayList<>(Arrays.asList(fixedArray)))){
            System.out.println("Running again, current array: " + Arrays.toString(fixedArray));
            fixedArray = repFix(fixedArray);
        }

        return fixedArray;
    }

    public Integer[] repFix(Integer[] arrayToFix){
        Integer[] fixedArray = arrayToFix;
        ArrayList<Integer> invalidIndexes = new ArrayList<>();

        // find invalids
        for(int i = 0; i < arrayToFix.length; i++){
            if(!isUpdateIndexValid(i, new ArrayList<>(Arrays.asList(fixedArray)))){
                invalidIndexes.add(i);

                int indexToSwapWith = getSwapIndex(i, arrayToFix);

                if(indexToSwapWith == -1){
                    continue;
                }

                int temp = fixedArray[indexToSwapWith];

                fixedArray[indexToSwapWith] = fixedArray[i];
                fixedArray[i] = temp;
            }
        }

        return fixedArray;
    }

    public int getSwapIndex(int toSwap, Integer[] arrayToFix){
        var rules = this.rules.get(arrayToFix[toSwap]);

        if(rules == null){
            return -1;
        }

        for(int i = 0; i < toSwap; i++){
            if(rules.contains(arrayToFix[i])){
                return i;
            }
        }

        return -1;
    }

    public boolean isMatchingRules(){
        return isMatchingRules(updates);
    }

    public boolean isMatchingRules(ArrayList<Integer> updates){
        boolean isValidUpdate = true;

        for(int currentUpdateIndex = 0; currentUpdateIndex < (long) updates.size(); currentUpdateIndex++){
            if(!isUpdateIndexValid(currentUpdateIndex, updates)){
                isValidUpdate = false;
            }
        }

        if(isValidUpdate){
            middleNumber = updates.get((updates.size() / 2));
        }

        return isValidUpdate;
    }

    public boolean isUpdateIndexValid(int index, ArrayList<Integer> updates){
        for(int validateIndex = index+1; validateIndex < (long)updates.size(); validateIndex++){
            if(rules.containsKey(updates.get(validateIndex))){
                var key = rules.get(updates.get(validateIndex));

                if(key.contains(updates.get(index))){
                    return false;
                }
            }
        }

        for(int validateIndex = index; validateIndex > -1; validateIndex--){
            if(Objects.equals(updates.get(validateIndex), updates.get(index))){
                continue;
            }

            if(rules.containsKey(updates.get(index))){
                var key = rules.get(updates.get(index));

                if(key.contains(updates.get(validateIndex))){
                    return false;
                }
            }
        }

        return true;
    }

    @Override
    public String toString() {
        return updates.toString();
    }
}
