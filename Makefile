TRAIN	= train
PREDICT	= predict 

SRC_TRAIN	 = train.go chart.go
TRAIN_PATH	 = ./training/
SRC_T = $(addprefix $(TRAIN_PATH), $(SRC_TRAIN:.go))

SRC_PREDICT	 = predict.go
PREDICT_PATH = ./prediction/
SRC_P = $(addprefix $(PREDICT_PATH), $(SRC_PREDICT:.go))

all: $(PREDICT) $(TRAIN)
	
$(TRAIN) : $(SRC_T)
	@go get ./...
	@go build -o $(TRAIN)

$(PREDICT) : $(SRC_P)
	@go get ./...
	@go build -o $(PREDICT)

clean:
	@rm -f $(TRAIN)
	@rm -f $(PREDICT)

re: clean all
