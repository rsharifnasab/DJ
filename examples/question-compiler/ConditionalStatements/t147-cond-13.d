
string str() {
    return "hey there!";
}

int main() {
    if (str() != "hey there!") Print("true");
    else Print("false");

    if (str() == "hey man!") Print("true");
    else Print("false");
}
