
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
    Item[] items;

    items = NewArray(2, Item);

    items[0] = new Item;
    items[0].init("Item #1");
    items[1] = new Item;
    items[1].init("Item #2");

    items[1].print();
    items[0].print();
}
