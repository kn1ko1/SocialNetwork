npx babel ./jsx/components/shared/*.jsx -d ./static/js/components/shared/
npx babel ./jsx/components/Chat/*.jsx -d ./static/js/components/Chat/
npx babel ./jsx/components/Groups/*.jsx -d ./static/js/components/Groups/
npx babel ./jsx/components/GroupDetail/*.jsx -d ./static/js/components/GroupDetail/
npx babel ./jsx/components/Home/*.jsx -d ./static/js/components/Home/
npx babel ./jsx/components/Notifications/*.jsx -d ./static/js/components/Notifications/
npx babel ./jsx/components/Profile/*.jsx -d ./static/js/components/Profile/

npx babel ./jsx/*.jsx -d ./static/js/

go run ./../Server/main.go