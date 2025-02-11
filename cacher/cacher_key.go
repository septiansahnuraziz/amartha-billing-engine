package cacher

import "amartha-billing-engine/utils"

func GetCustomerCacheKeyByID(customerID uint64) string {
	return createCacheKey(utils.WriteStringTemplate("cache:object:customer:id:%d", customerID))
}

func GetMasterDealerChannelCacheKeyByID(channelId uint64) string {
	return createCacheKey(utils.WriteStringTemplate("cache:object:dealer_channel:id:%d", channelId))
}

func GetDealerStoreCacheKeyByID(storeId uint64) string {
	return createCacheKey(utils.WriteStringTemplate("cache:object:dealer_store:id:%d", storeId))
}

func GetMultipleDealerStoreCacheKeyByID(storeIds []uint64) (keys []string) {
	for _, storeId := range storeIds {
		keys = append(keys, createCacheKey(utils.WriteStringTemplate("cache:object:dealer_store:id:%d", storeId)))
	}
	return
}

func GetMultipleMasterDealerChannelCacheKeyByID(channelIds []uint64) (keys []string) {
	for _, channelId := range channelIds {
		keys = append(keys, createCacheKey(utils.WriteStringTemplate("cache:object:dealer_channel:id:%d", channelId)))
	}
	return
}

func GetMasterProductKeyByImei(imei string) string {
	return createCacheKey(utils.WriteStringTemplate("cache:object:master_product:imei:%d", imei))
}

func GetMultipleMasterProductKeyByImei(imeiList []string) (keys []string) {
	for _, masterImei := range imeiList {
		keys = append(keys, createCacheKey(utils.WriteStringTemplate("cache:object:master_product:imei:%d", masterImei)))
	}
	return
}

func GetMultipleMdrBankKeyById(bankIdList []string) (keys []string) {
	for _, id := range bankIdList {
		keys = append(keys, GetMdrBankCacheKeyByID(utils.ExpectedNumber[uint64](id)))
	}
	return
}

func GetOrderCacheKeyByID(orderId uint64) string {
	return createCacheKey(utils.WriteStringTemplate("cache:object:order:id:%d", orderId))
}

func GetMdrBankCacheKeyByID(id uint64) string {
	return createCacheKey(utils.WriteStringTemplate("cache:object:mdr_bank:id:%d", id))
}

func GetWebSocketMessageCacheKey(messageID string) string {
	return createCacheKey(utils.WriteStringTemplate("cache:object:websocket:message:%s", messageID))
}

func GetWebSocketUserMessagesCacheKey(userEmail string) string {
	return createCacheKey(utils.WriteStringTemplate("cache:object:websocket:user:%s:messages", userEmail))
}

func GetWebSocketPendingMessagesCacheKey() string {
	return createCacheKey("cache:object:websocket:pending_messages")
}

func GetWebSocketAckMessageCacheKey(messageID string) string {
	return createCacheKey(utils.WriteStringTemplate("cache:object:websocket:ack:%s", messageID))
}

func GetUserScopeCacheKey(email string) string {
	return createCacheKey(utils.WriteStringTemplate("cache:object:user:scope:%s", email))
}

func GetDealerCacheKeyByID(id uint64) string {
	return createCacheKey(utils.WriteStringTemplate("cache:object:dealer:id:%d", id))
}

func GetFeeConfigurationCacheKeyByID(id uint64) string {
	return createCacheKey(utils.WriteStringTemplate("cache:object:fee_configuration:id:%d", id))
}

func GetUserPrincipalCacheKey(key string) string {
	return createCacheKey(utils.WriteStringTemplate("cache:object:user:principal:%s", key))
}

func LockCreateOrderByOrderNumberLockKey(orderNumber string) string {
	return createCacheKey(utils.WriteStringTemplate("cache:create_order:order_number:%s", orderNumber))
}

func LockIndexDealerReportingByOrderNumberLockKey(orderNumber string) string {
	return createCacheKey(utils.WriteStringTemplate("cache:index_dealer_reporting:order_number:%s", orderNumber))
}

func LockIndexFinanceReportingByOrderNumberLockKey(orderNumber string) string {
	return createCacheKey(utils.WriteStringTemplate("cache:index_finance_reporting:order_number:%s", orderNumber))
}

func LockPushPaymentByOrderIdLockKey(orderId string) string {
	return createCacheKey(utils.WriteStringTemplate("cache:push_payment:order_id:%s", orderId))
}
