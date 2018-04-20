NAME	=	ft_linear_regression 
SRCS	=	$(shell find . -type f -name "*.go")

all: $(NAME)
	
$(NAME)	: $(SRCS)
	@go get ./...
	@go build -o $(NAME)
clean:
	@rm -f $(NAME)
re: clean all
