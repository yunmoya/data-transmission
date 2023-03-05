from blspy import PrivateKey, SignatureMPL, PublicKeyMPL, BasicSchemeMPL
from random import SystemRandom


class BLS:

    @staticmethod
    def createKey():
        seed = []
        for i in range(0, 32):
            seed.append(SystemRandom().randrange(0, 254))
        seed = bytes(seed)
        private_key = BasicSchemeMPL.key_gen(seed)
        public_key = private_key.get_g1()
        return private_key, public_key

    # @staticmethod
    # def serialize(_obj):
    #     try:
    #         if (type(_obj) == PrivateKey) or (type(_obj) == PublicKey) or (type(_obj) == Signature):
    #             return str(_obj.serialize(), "ISO-8859-1")
    #         else:
    #             _obj
    #     except Exception as e:
    #         print("given parameters compatibility error")
    #         print(e)
    #         return None

    @staticmethod
    def deserialize(_obj, _type):
        try:
            if (type(_obj) != PrivateKey) or (type(_obj) != PublicKeyMPL) or (type(_obj) != SignatureMPL):
                _obj = bytes(_obj, "ISO-8859-1")
                if _type == PrivateKey:
                    return PrivateKey.from_bytes(_obj)
                elif _type == PublicKeyMPL:
                    return PublicKeyMPL.from_bytes(_obj)
                elif _type == SignatureMPL:
                    return SignatureMPL.from_bytes(_obj)
            else:
                _obj
        except Exception as e:
            print("given parameters compatibility error")
            print(e)
            return None

    @staticmethod
    def sign_data(_data, private_key):

        if type(private_key) != PrivateKey:
            private_key = BLS.deserialize(private_key, PrivateKey)
            if type(private_key) != PrivateKey:
                return None
        if type(_data) != str:
            return None
        temp_array = []
        for c in _data:
            temp_array.append(ord(c))
        msg = bytes(temp_array)

        _sig = private_key.sign(msg)
        return _sig
        # return str(sig.serialize(), "ISO-8859-1")

    @staticmethod
    def combine_signature(_signatures):
        try:
            temp_list = []
            if type(_signatures) == list:
                print("check_1")
                for _sign in _signatures:
                    if type(_sign) == SignatureMPL:
                        print("check_2")
                        # _sign = copy(_sign)
                        temp_list.append(_sign)
                    else:
                        print("check_3")
                        _sign = BLS.deserialize(_sign, SignatureMPL)
                        if type(_sign) == SignatureMPL:
                            temp_list.append(_sign)
                        else:
                            return None
                print("exit for")
            else:
                if type(_signatures) == SignatureMPL:
                    print("check_4")
                    temp_list.append(_signatures)
                else:
                    _sign = BLS.deserialize(_signatures, SignatureMPL)
                    print("check_5")
                    if type(_sign) == SignatureMPL:
                        temp_list.append(_sign)
                    else:
                        return None
            print("antim padav")
            return SignatureMPL.aggregate(temp_list)
        except Exception as e:
            print("given parameters compatibility error")
            print(e)
            return None

    # @staticmethod
    # def verify_sign(_data, _signature, public_key_list):
    #
    #     if type(_signature) != SignatureMPL:
    #         _signature = BLS.deserialize(_signature, SignatureMPL)
    #         if type(_signature) != SignatureMPL:
    #             return None
    #
    #     if type(_data) != str:
    #         return None
    #     temp_array = []
    #     for c in _data:
    #         temp_array.append(ord(c))
    #     msg = bytes(temp_array)
    #
    #     agg_info_list = []
    #
    #     if type(public_key_list) == list:
    #         for pk in public_key_list:
    #             if type(pk) == PublicKeyMPL:
    #                 agg_info = AggregationInfo.from_msg(pk, msg)
    #                 agg_info_list.append(agg_info)
    #             else:
    #                 pk = BLS.deserialize(pk, PublicKey)
    #                 if type(pk) == PublicKey:
    #                     agg_info = AggregationInfo.from_msg(pk, msg)
    #                     agg_info_list.append(agg_info)
    #                 else:
    #                     return None
    #     else:
    #         if type(public_key_list) == PrivateKey:
    #             agg_info = AggregationInfo.from_msg(public_key_list, msg)
    #             agg_info_list.append(agg_info)
    #         else:
    #             public_key_list = BLS.deserialize(public_key_list, PublicKey)
    #             if type(public_key_list) == PublicKey:
    #                 agg_info = AggregationInfo.from_msg(public_key_list, msg)
    #                 agg_info_list.append(agg_info)
    #             else:
    #                 return None
    #     agg_public_key = AggregationInfo.merge_infos(agg_info_list)
    #     _signature.set_aggregation_info(agg_public_key)
    #     return _signature.verify()


# sk1 = BLS.createKey()
# pk1 = sk1.get_public_key()
# sk2 = BLS.createKey()
# pk2 = sk1.get_public_key()
# sk3 = BLS.createKey()
# pk3 = sk1.get_public_key()
#
# data_to_sign = "hello stranger, you are going used for signing";
# sig_1 = BLS.sign_data(data_to_sign, sk1)
# sig_2 = BLS.sign_data(data_to_sign, sk2)
# sig_3 = BLS.sign_data(data_to_sign, sk3)
# # sig_1 = BLS.deserialize(sig_1, Signature)
#
# by = sig_1.serialize()
# # sig_11 = str(sig_1.serialize(), "ISO-8859-1")
# sig = Signature.from_bytes(by)
# if type(sig_1) == type(sig):
#     if sig == sig_1:
#         print("ckdjvod")
# print(sig)
# print(sig_1)
#
# cosig = Signature.aggregate([sig_2, sig])
# # cosig = BLS.combine_signature(sig_1)
# print(cosig)
# # verify
