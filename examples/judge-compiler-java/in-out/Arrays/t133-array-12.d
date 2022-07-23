
class Item {
    string content;

    void init(string s) {
        content = s;
    }

    void print() {
        Print(content);
    }
}

int main() {
    Item[][] items;

    items = NewArray(1, Item[]);
    items[0] = NewArray(1, Item);

    items[0][0] = new Item;
    items[0][0].init("The Only Item!");

    items[0][0].print();
}
