class Person {

    void set_name(string name) {
        this.name = name;
    }
    string name;

    string get_desc() {
        Print("my name is ", name);
    }
}

class Student extends Person {
    int grade;
    string get_desc() {
        Print("I am ", name, " , A student in grade : ", this.grade);
    }

    void set_grade(int grade) {
        this.grade = grade;
    }

}

class Graduate extends Student {
    void set_year (int y ) {
        year = y;
    }
    int year;
    string get_desc()
    {
        Print("I am Grade ", grade);
        Print(this.name);
        Print(year);
        Print("I am ", name, " a Graduated student in year ", year);
    }
}

class UnderGraduate extends Student {
    int year;
    void set_year (int y) {
        year = y;
    }

    string get_desc() {
        Print("I am Grade ", grade);
        Print("I am ", name, ", an undergraduate student which will graduate at ", year);
    }

}

int main() {
    string type;
    string name;
    Person p;
    Student s;
    Graduate g;
    UnderGraduate ug;

    type = ReadLine();
    if(type == "G") {
        int grade;
        Print("Name, Grade, Year Of Graduation");
        g = new Graduate;
        name = ReadLine();
        g.set_name(name);
        g.set_grade(ReadInteger());
        g.set_year(ReadInteger());
        p = g;
        s = g;
    }
    else {
    int grade;
        Print("Name, Grade, Expected Year Of Graduation");
        ug = new UnderGraduate;
        name = ReadLine();
        ug.set_name(name);
        ug.set_grade(ReadInteger());
        ug.set_year(ReadInteger());
        p = ug;
        s = ug;
    }

    p.get_desc();
    s.get_desc();

}
