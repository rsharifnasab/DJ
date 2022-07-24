import java.io.File;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collection;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Judger {

    public static String getTestName(File f) {
        String filename = f.getName();

        if(filename.indexOf(".") > 0) {
            filename = filename.substring(0, filename.lastIndexOf("."));
        }

        return filename;
    }

    static String getTestsFolder(){
        return "testgroup";
    }

    public static List<File> allTestsFiles(String suffix) throws Exception {
        final File testFolder = new File(getTestsFolder());
        //System.out.println(testFolder);
        final List<File> tests = Arrays.stream(testFolder.listFiles())
            .map(a -> a.toString())
            .sorted()
            .filter(a -> a.matches(".*\\." + suffix))
            .map(s -> new File(s))
            .collect(Collectors.toCollection(ArrayList::new));
        return tests;
    }

    public static File aTestFile(int i, String suffix) throws Exception {
        return allTestsFiles(suffix).get(i);
    }

    public static void runTest(int i) throws Exception {
        File inpFile = aTestFile(i, "d");
        File actualFile = new File("compiled.s");
        File expected = aTestFile(i, "out");
        System.out.println(expected.getAbsoluteFile().toString());
        //compiler.Main.run(inpFile, actualFile);

    }

    public static int countTests() throws Exception {
        return allTestsFiles("d").size();
    }


    public static void main(String[] args) throws Exception {
        final String command = args[0];

        switch (command) {
            case "count":
                System.out.println(countTests());
                break;

            case "test":
                runTest(Integer.parseInt(args[1]));
                break;

            default:
                throw new Error("bad arguments : " + Arrays.toString(args));
        }
    }

}
