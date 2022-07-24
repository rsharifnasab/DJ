
int main() {
    if (true) {
        if (true || !true) {
            Print("a");
        }
        else {}
    } else {
        if (false) {
            Print("not gonna happen!");
        }
    }

    if (false) {
        if (true || !true) {
            Print("a");
        }
        else {}
    } else {
        if (true) {
            Print("this may happen!");
        }
    }
}
