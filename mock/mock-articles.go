package mock

import (
	"fmt"
	"math/rand"
	"thelight/models"
)

func returnRandomArticleID() models.Article {

	randID := rand.Int31()

	return models.Article{
		ID:       fmt.Sprintf("%v", randID),
		Date:     "17th May 2021",
		Title:    "What is Cryptocurrency",
		Preview:  "Cryptocurrency is decentralized digital money, based on blockchain technology. You may be familiar with the most popular versions, Bitcoin and Ethereum, but there are more than 5,000 different cryptocurrencies in circulation, according to CoinLore. You can use crypto to buy regular goods and services, although many people invest in cryptocurrencies as they would in other assets, like stocks or precious metals. While cryptocurrency is a novel and exciting asset class, purchasing it can be risky as you must take on a fair amount of research to fully understand how each system works.",
		Body:     "Cryptocurrency is decentralized digital money, based on blockchain technology. You may be familiar with the most popular versions, Bitcoin and Ethereum, but there are more than 5,000 different cryptocurrencies in circulation, according to CoinLore. You can use crypto to buy regular goods and services, although many people invest in cryptocurrencies as they would in other assets, like stocks or precious metals. While cryptocurrency is a novel and exciting asset class, purchasing it can be risky as you must take on a fair amount of research to fully understand how each system works. Cryptocurrency is decentralized digital money, based on blockchain technology. You may be familiar with the most popular versions, Bitcoin and Ethereum, but there are more than 5,000 different cryptocurrencies in circulation, according to CoinLore. You can use crypto to buy regular goods and services, although many people invest in cryptocurrencies as they would in other assets, like stocks or precious metals. While cryptocurrency is a novel and exciting asset class, purchasing it can be risky as you must take on a fair amount of research to fully understand how each system works. Cryptocurrency is decentralized digital money, based on blockchain technology. You may be familiar with the most popular versions, Bitcoin and Ethereum, but there are more than 5,000 different cryptocurrencies in circulation, according to CoinLore. You can use crypto to buy regular goods and services, although many people invest in cryptocurrencies as they would in other assets, like stocks or precious metals. While cryptocurrency is a novel and exciting asset class, purchasing it can be risky as you must take on a fair amount of research to fully understand how each system works. Cryptocurrency is decentralized digital money, based on blockchain technology. You may be familiar with the most popular versions, Bitcoin and Ethereum, but there are more than 5,000 different cryptocurrencies in circulation, according to CoinLore. You can use crypto to buy regular goods and services, although many people invest in cryptocurrencies as they would in other assets, like stocks or precious metals. While cryptocurrency is a novel and exciting asset class, purchasing it can be risky as you must take on a fair amount of research to fully understand how each system works. Cryptocurrency is decentralized digital money, based on blockchain technology. You may be familiar with the most popular versions, Bitcoin and Ethereum, but there are more than 5,000 different cryptocurrencies in circulation, according to CoinLore. You can use crypto to buy regular goods and services, although many people invest in cryptocurrencies as they would in other assets, like stocks or precious metals. While cryptocurrency is a novel and exciting asset class, purchasing it can be risky as you must take on a fair amount of research to fully understand how each system works. Cryptocurrency is decentralized digital money, based on blockchain technology. You may be familiar with the most popular versions, Bitcoin and Ethereum, but there are more than 5,000 different cryptocurrencies in circulation, according to CoinLore. You can use crypto to buy regular goods and services, although many people invest in cryptocurrencies as they would in other assets, like stocks or precious metals. While cryptocurrency is a novel and exciting asset class, purchasing it can be risky as you must take on a fair amount of research to fully understand how each system works.",
		Tag:      []string{"Finance", "Technology", "Cryptocurrency", "Knowledge"},
		ImageURL: "https://thumbor.forbes.com/thumbor/fit-in/900x510/https://www.forbes.com/advisor/wp-content/uploads/2020/11/what-is-crypto.jpg",
		WriterInfo: models.WriterInfo{
			ID:        "1a",
			Name:      "Hadekha",
			AvatarURL: "https://pbs.twimg.com/profile_images/1363210545118150659/Uo-XiGtv_400x400.jpg",
			Bio:       "I am a fullstack developer",
		},
	}
}

var Onearticle models.Article = returnRandomArticleID()

var Articles []models.Article = []models.Article{
	returnRandomArticleID(),
	returnRandomArticleID(),
	returnRandomArticleID(),
	returnRandomArticleID(),
	returnRandomArticleID(),
	returnRandomArticleID(),
}
