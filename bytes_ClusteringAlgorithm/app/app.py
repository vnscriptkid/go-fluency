from flask import Flask, render_template, request
import pandas as pd
import re
from sklearn.cluster import KMeans
from sklearn.feature_extraction.text import TfidfVectorizer

app = Flask(__name__)

@app.route('/', methods=['GET', 'POST'])
def index():
    if request.method == 'POST':
        # Get the uploaded file
        file = request.files['file']

        # Read the file as a DataFrame
        data = pd.read_csv(file)

        # Preprocess the text data
        def preprocess(text):
            # Removes any URLs in the text by replacing them with an empty string
            text = re.sub(r"http\S+", "", text)
            # Removes any non-alphanumeric characters from the text by replacing them with an empty string
            text = re.sub(r"[^a-zA-Z0-9\s]", "", text)
            # Converts all the text to lowercase
            text = text.lower()
            return text
        
        # Selects the text column of the data DataFrame
        # applies `preprocess` to each element in a column or row of a DataFrame
        data['text'] = data['text'].apply(preprocess)

        # Extract features using TF-IDF vectorizer
        # stop words (e.g., "the", "and", "a") should be removed from the text before vectorization
        tfidf_vectorizer = TfidfVectorizer(stop_words='english')
        tfidf = tfidf_vectorizer.fit_transform(data['text'])
        #             term_1  term_2  term_3  term_4  ...  term_n
        # tweet_1     0.01    0.00    0.02    0.00         0.00
        # tweet_2     0.00    0.03    0.01    0.00         0.00
        # tweet_3     0.02    0.01    0.00    0.00         0.00
        # ...         ...     ...     ...     ...          ...
        # tweet_m     0.00    0.00    0.00    0.04         0.02


        # Perform text clustering using K-means algorithm
        kmeans = KMeans(n_clusters=5, init='k-means++', max_iter=100, n_init=1)
        kmeans.fit(tfidf)

        # Get the labels for each tweet
        labels = kmeans.labels_

        # Add the labels to the DataFrame
        data['label'] = labels

        # Render the template with the clustered data
        return render_template('index.html', data=data.to_html())
        
    # Render the form template if no file has been uploaded
    return render_template('form.html')

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')
