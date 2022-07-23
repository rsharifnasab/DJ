
class Base {
    int counter;
    Base action_s() {}
    Base action_o() {}
    void print() {
        Print(counter);
    }
}

class A extends Base {
    B pair;

    Base action_s() {
        Print("A.action_s()");
        counter = counter + 1;
        return this;
    }

    Base action_o() {
        Print("A.action_o()");
        counter = counter + 1;
        return pair;
    }

    void set_pair(B p) {
        pair = p;
    }
}

class B extends Base {
    A pair;

    Base action_s() {
        Print("B.action_s()");
        counter = counter + 1;
        return this;
    }

    Base action_o() {
        Print("B.action_o()");
        counter = counter + 1;
        return pair;
    }

    void set_pair(A p) {
        pair = p;
    }
}

int main() {
    Base base;
    A a;
    B b;
    a = new A;
    b = new B;
    a.set_pair(b);
    b.set_pair(a);

    base = a;
    base.action_s().action_s().action_o().action_o().action_s().action_o().action_s().action_s().action_o();
    a.print();
    b.print();
}