npx babel jsx/shared/*.jsx -d static/js/shared/
npx babel jsx/components/*.jsx -d static/js/components/
npx babel jsx/*.jsx -d static/js/

go run .