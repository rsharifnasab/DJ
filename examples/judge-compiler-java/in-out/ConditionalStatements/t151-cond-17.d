
int main() {
    if (true) {
        if (true || !true) {
            Print("a");
        }
    }

    if (false) {
        if (true || !true) {
            Print("b");
        }
    }
}
