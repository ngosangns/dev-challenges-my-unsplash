package database

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/ngosangns/devchallenges-my-unsplash-api/flags"
	"google.golang.org/api/option"
)

const creds = `{
	"type": "service_account",
	"project_id": "ngosangns-myunsplash",
	"private_key_id": "5d8e3e4f705686f674149632e536e467e60e859f",
	"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDeSlyxBhVJnpTf\nF7uQUeydxTZAK5nnhYDOotXvUtTPDc1fxoPMgB0NI+GicF9HYswkGYoqFb38BnVz\n3jx5mBxy5+JLXSgywvK5+kE6jKE3g4SkOHhd8GpHl8oq1RPJ/fk5iw4H0ykscsDM\nGalXbNtP797P0plFPZbpMufB8Ktf+NcaTCAKp8kV1mH4wYYdQc3/4W0P0miE3wR6\nSnpY+41zH5m77J3kKAbsHVAmggliaC6IsdLy/sP0kSWmBpmPx82ZpOSweWqEugbI\nD0XSf5CerO6W5orCoeI3GmuzPaq/gQ6ubhBSkULMDib0b5lZuYAr/vZ/vZVh+CKA\n/0HLBR4FAgMBAAECggEAF2IfKggHyTzDUEyA0IetqlMKh+mOQQqMKDB87LX76OP8\nzKxnROXdKHqzy5CilTHRY6LEI/UxjQdo541jMy9LHtNBvMfJ6L0VJ3cp0WzlsJ9H\nHrdKAie+JlXTUHoloMP0qZb1HOGbnS9dFJtGcwnicCA9ggMXWG9IJrPSKVaV0Hx8\ngeV2R8Y6Z2sZzBxb3FCPiQTQDZVOt6PHRiVs3uSOf5dmK1PkAtW22DWAtAmBSfLw\n/EhX7fFtWJMr529/n/LrWSkjnXkh0Px9Knf+WQGKyt5/S4c7Zv2jO8G6B/K7FsRx\nsWWiu8b3HHyKCVm3pihjBmZgA5qhg7crjzJlyH1lKQKBgQDy4pwzSU+rV+lTeaUo\nrPCodfB1UOiYHeqdlkfaTUoi5/QB8wQQlWwvEx30b/2ZlLVTP2eC9A2en1p1C2J3\nd7o9AYAy2yHv5OFFCtdOz1qGnH8kC7Wd6XvRJVpUwkZj44480k7G1s7SLjAbxc/M\nJqq8xEt7Jo+y3jn/qKr9mOD8OwKBgQDqSxJ4oRfeQZ7h0RinFkoFxJmD37Bk5uYu\n0k8CNykocKkdJGdZ4zRkJEq0s9YB6G0os6qOJajC7aphfJXCm0NvUcs/cE9pBMPO\nxk96NK9oKhY+B1jFs3TbLHN0FbM1ks/4oiTtMNW3afelg5hpdn2/z5+jFQ1Qaul2\nAs/PoyHqvwKBgQDPi0ZQKMf0MRXG575yaHzoZodHqBBo07+EU0MozjyYYT5HYCJM\nDXK5Mwu4sO2yC7l53kcURnlQ3vhL5aOJVHyG3r5uNoPkWt8HwUDELjk7p4SZB4x0\n2GueBt4OaUf+2RbU8ByymiF2xlMDmF99IllpHg8lt4i5Hl3PEZCzTKyLuwKBgQCJ\nPzbv0TaBLSGcmPd/EbCFFnjzZnsXcKdQsEQ2QQfsN6rpdmTQpCcRrqzQlDR+Y6Bm\nYnE72dW1crIHb4rvODLVqC2O2QS7uV3W6BawuOIz+8uRxI39KjX/Glf7naagcINB\n3AITrDLVwc6fd6o1xYK+dSrHxtJI1Ig0aEtCEAcTUwKBgDz1emR9VVmPFGNf/qb0\nbWUTdAaRPkobVOGtPUl1vb+fytcXQ4bT/7rF6DMlefHAxcJ2KiwLA8TgQ7wOmKgh\nXGtiAUZjIrTZutDI26CmXI+PF/8bjRdEm1+M1gyrROjjf6KQ+6T5sZVQgYnIB2EY\nstPrU8oiGDj3+biWILcXnYwL\n-----END PRIVATE KEY-----\n",
	"client_email": "firebase-adminsdk-zmbu3@ngosangns-myunsplash.iam.gserviceaccount.com",
	"client_id": "112304039689025329319",
	"auth_uri": "https://accounts.google.com/o/oauth2/auth",
	"token_uri": "https://oauth2.googleapis.com/token",
	"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-zmbu3%40ngosangns-myunsplash.iam.gserviceaccount.com"
}`

func Connect() (*firestore.Client, context.Context, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, flags.DbName.Get(), option.WithCredentialsJSON([]byte(creds)))

	// defer client.Close()
	if err != nil {
		return client, ctx, err
	}
	return client, ctx, nil
}
