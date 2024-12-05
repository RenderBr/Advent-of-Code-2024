import java.io.File;
import java.io.FileNotFoundException;
import java.util.*;

public class Main {
    public static HashMap<Integer, ArrayList<Integer>> ruleMap = new HashMap<>();
    public static ArrayList<Update> updates = new ArrayList<>();

    public static void main(String[] args) {
        load();

        int sum = 0;
        int fixedSum = 0;
        for(Update update : updates){

            if(update.isMatchingRules()){
                sum += update.middleNumber;
            }else{
                System.out.println(update);

                var fixed = update.fix();

                System.out.println(Arrays.toString(fixed));

                fixedSum += fixed[fixed.length / 2];
            }
        }

        System.out.println("Sum: " + sum);
        System.out.println("FixedSum: " + fixedSum);

    }

    public static void load() {
        try {
            boolean rules = true;
            File dataFile = new File("data.txt");
            Scanner fileReader = new Scanner(dataFile);

            while (fileReader.hasNextLine()) {
                String line = fileReader.nextLine();

                if (line.isEmpty()) {
                    rules = false;
                    continue;
                }

                if (rules) {
                    var rulings = line.split("\\|");

                    if (rulings.length == 2) {
                        int rulingNumber = Integer.parseInt(rulings[0]);
                        int otherNumber = Integer.parseInt(rulings[1]);

                        if (ruleMap.containsKey(rulingNumber)) {
                            ruleMap.get(rulingNumber).add(otherNumber);
                        } else {
                            var newList = new ArrayList<Integer>();
                            newList.add(otherNumber);
                            ruleMap.put(rulingNumber, newList);
                        }
                    }


                } else {
                    var updateLine = line.split(",");
                    Update updateObj = new Update(ruleMap);

                    for (String update : updateLine) {
                        updateObj.updates.add(Integer.parseInt(update));
                    }

                    updates.add(updateObj);
                }

            }
        } catch (FileNotFoundException e) {
            throw new RuntimeException(e);
        }
    }
}