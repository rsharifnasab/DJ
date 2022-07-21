
int main() {
    if (true) {
        if (true || !true)
            Print("a");
        else
			Print("bb");
    } else {
        if (false)
            Print("not gonna happen!");
    }
}
