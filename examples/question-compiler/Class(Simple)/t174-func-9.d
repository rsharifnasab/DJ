
class Item {
    string content;

    void init(string s) {
        content = s;
    }

    void print() {
        Print(content);
    }
}

Item test() {
    Item i;
    i = new Item;
    i.init("Special Item!");
    return i;
}

int main() {
    Item i;
    i = test();
    i.print();
}
