<!--
 * @Author: Caspar
 * @Date: 2022-07-29 20:48:16
 * @Description: file content
-->
#����
![Alt text](https://github.com/gochenzl/chess/blob/master/doc/pic/architecture.png?raw=true)
��Ϊ����㡢�߼��㡢���ݲ㣬redis��server_centerΪЭ���㡣 

һ��server_gate���Դ�һ�����߶��server_game��ÿ��server_gate����һ��Ψһ��gateid����server_gate������server_gameʱ�����gateid���͸�server_game��  

ÿ���ͻ��˵�����������gateid��connid(����id)��ʶ���ͻ��˵�¼�ɹ�����server_game���浽server_center������server_center�ַ������е�server_game��������ÿ��server_game�������пͻ��˵�������Ϣ��һ�����棬������������Ϣ��ת����Ϣ������ͻ������ӵ�gateid��server_game����ͬ����ֱ��ת����server_gate�������һ��������뵽��Ӧ��redis����Ϣ�����У���Ӧ��server_gate��ȡ����ת����  

���ݲ����server_table��user_db��  
user_db���������Ϣ������Ϊ�κ�֧��redisЭ������ݿ⡣  
server_table�������������߼���������������ѯ�ͱ���������Ϣ����ѯ��ҵķ���λ�á������ɹ�ʱ����redis������Ϣ��server_game��ȡ�������������г�ʱ��Ϣʱ����redis������Ϣ��server_game��ȡ��������  

server_login(ͼ��û�л���)�����˺ŵ�¼����΢�ŵ�¼�ȡ��˺ŵ�¼�ɹ�����redis������ҵ�¼�ɹ��ļ�¼�������ͻ��˷�����Ϸ��¼��ַ�͵�¼token��  
server_login��server_game֮��ͨ��redisͨ�ţ�����֤token����ȡ�����Ϣ�����Ը��ݿͻ��˰汾ѡ����Ϸ��¼��ַ��������¡�

##server_gate
* �ͻ����������ӹ���
* ת���ͻ�����Ϣ���߼���
* ת���߼�����Ϣ���ͻ���
* ���͹㲥��Ϣ

##server_game
* ����ҵ���߼�
* ά���ͻ���������Ϣ(�߼����г�Ϊsession)
* ת����Ӧ��Ϣ�������

##server_table
* ���뷿�䡢�뿪���䡢����
* ��ѯ�͸���������Ϣ
* ��ѯ��ҵķ���λ��

##server_center
* �ͻ���������Ϣ����
* �����ӡ�ɾ��������Ϣʱ���ַ����߼���
* ������Ϣ�־û�������������ʧ

#��չ��
* �������߼������������չ
* server_centerֻ����������Ϣ������������ɾ�ﵽÿ��3������ң������Ϊƿ��
* server_table��Ҫ��������������Ϣ�Ĳ�ѯ�͸���(�ֱ���ԣ�ÿ��ɴ�8w������)������ﵽƿ��������ÿ�����俪һ��(��������£���ѯ������ĸ�����Ҫ��ѯ���е�server_table)
* user_db�Ƽ�ʹ�þ��г־û����ܵ�ssdb�����ݴﵽһ����ģʱ����Ȼ���нϸߵ����ܣ���չ���Ը���userid����ϣ������Ҳ����ʹ��LedisDB��