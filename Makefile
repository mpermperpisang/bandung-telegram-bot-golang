vendor-prepare:
  @go get -u gopkg.in/tucnak/telebot.v2
  @go get github.com/joho/godotenv
  @go get gopkg.in/Iwark/spreadsheet.v2
  @go get -u github.com/go-sql-driver/mysql
  @echo "Package installed"
